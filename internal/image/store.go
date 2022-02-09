package image

import (
	"docker-/internal/constants"
	"docker-/pkg/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

func ReadAll(excludeImage *Image) ([]*Image, error) {

	if err := utils.PathInit(constants.DefaultImageConfigPath); err != nil {
		return nil, fmt.Errorf("image repositories file path init eroor %v", err)
	}

	// read repositories
	bytes, err := ioutil.ReadFile(constants.DefaultImageConfigPath)
	if err != nil {
		return nil, fmt.Errorf("read images repositories file %s, error %v", constants.DefaultImageConfigPath, err)
	}

	// empty file
	if len(bytes) == 0 {
		return nil, nil
	}

	// decode
	var images []*Image
	if err := json.Unmarshal(bytes, &images); err != nil {
		return nil, fmt.Errorf("json decode error %v", err)
	}

	// exclude image
	if excludeImage != nil && len(images) > 0 {
		if index, err := index(images, excludeImage); err == nil {
			images = append(images[:index], images[index+1:]...)
		}
	}

	return images, nil
}

func Save(image *Image) error {

	images, err := ReadAll(nil)
	if err != nil {
		return err
	}

	if err := writeImage(append(images, image)); err != nil {
		return err
	}
	return nil
}

func Remove(image *Image) error {

	images, err := ReadAll(image)
	if err != nil {
		return err
	}

	// write file
	if err := writeImage(images); err != nil {
		return err
	}
	return nil
}

func Update(image *Image) error {

	images, err := ReadAll(image)
	if err != nil {
		return err
	}

	// write file
	if err := writeImage(append(images, image)); err != nil {
		return err
	}
	return nil
}

func GetImageByNameOrId(identifier string) *Image {
	images, err := ReadAll(nil)
	if err != nil {
		return nil
	}
	for _, image := range images {
		if identifier == image.Id || identifier == image.Name {
			return image
		}
	}
	return nil
}

func GetDockerInspectInfo(name string) (*Image, error) {

	format := strings.Join(constants.InspectFormatArgs, "|")
	cmd := exec.Command("docker", "inspect", "-f", format, name)
	bytes, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// notes: need to remove the tailing newline \n.
	info := strings.Split(strings.Trim(string(bytes), "\n"), "|")
	repoTag := strings.Split(name, ":")
	image := &Image{
		Id:         info[0][7:],
		Name:       repoTag[0],
		Tag:        repoTag[1],
		Size:       info[1],
		Counts:     0,
		Work:       info[2],
		Hostname:   info[3],
		Entrypoint: []string{},
		Command:    []string{},
		Envs:       []string{},
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	utils.Unmarshal([]byte(info[4]), &image.Entrypoint)
	utils.Unmarshal([]byte(info[5]), &image.Command)
	utils.Unmarshal([]byte(info[6]), &image.Envs)

	return image, nil
}

func writeImage(images []*Image) error {
	bytes, err := json.Marshal(images)
	if err != nil {
		return fmt.Errorf("images json encode error %v", err)
	}

	if err := ioutil.WriteFile(constants.DefaultImageConfigPath, bytes, 0644); err != nil {
		return fmt.Errorf("write images to file error %v", err)
	}

	return nil
}

func index(images []*Image, image *Image) (int, error) {
	for i, img := range images {
		if img.Id == image.Id || img.Name == image.Name {
			return i, nil
		}
	}
	return 0, fmt.Errorf("search image not match")
}
