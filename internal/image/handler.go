package image

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"
)

func PullHandler(name string) error {

	// check name
	if !strings.Contains(name, ":") {
		name = name + ":latest"
	}

	// check exist
	if image := GetImageByNameOrId(name); image != nil {
		return fmt.Errorf("%s already exist", name)
	}

	// check docker image , not exist, download image
	image, err := GetDockerInspectInfo(name)
	if err != nil {

		// download image
		cmd := exec.Command("docker", "pull", name)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if er := cmd.Run(); er != nil {
			return er
		}

		// inspect image info
		img, e := GetDockerInspectInfo(name)
		if e != nil {
			return e
		}

		image = img
	}

	return Save(image)
}

func ListHandler() error {

	images, err := ReadAll(nil)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 12, 1, 3, ' ', 0)
	_, _ = fmt.Fprint(w, "IMAGE ID\tNAME\tTAG\tCOUNTS\tSIZE\tCREATED\n")

	for _, image := range images {
		fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%s\t%s\n",
			image.Id,
			image.Name,
			image.Tag,
			image.Counts,
			image.Size,
			image.CreateTime,
		)
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("failed to flush %v", err)
	}
	return nil
}

func RemoveHandler(identifier string) error {

	// check exist
	image := GetImageByNameOrId(identifier)
	if image == nil {
		return fmt.Errorf("%s not exist", identifier)
	}

	// container use image count
	if image.Counts > 0 {
		return fmt.Errorf("remove images failed, there still exist %d containers using the image %s", image.Counts, identifier)
	}

	return Remove(image)
}

func ChangeCounts(identifier, action string) error {

	image := GetImageByNameOrId(identifier)
	if image == nil {
		return fmt.Errorf("image not exist")
	}

	switch action {
	case "create":
		image.Counts++
	case "delete":
		image.Counts--
	}

	return Update(image)
}
