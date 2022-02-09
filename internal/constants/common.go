package constants

const DefaultInstallDir = "/var/lib/docker-"

const DefaultContainerDir = DefaultInstallDir + "/containers"

const DefaultNetworkDir = DefaultInstallDir + "/networks"

const DefaultImagesDir = DefaultInstallDir + "/images/"

const DefaultImageConfigPath = DefaultImagesDir + "repositories.json"

var InspectFormatArgs = []string{
	"{{.Id}}",
	"{{.Size}}",
	"{{.ContainerConfig.WorkingDir}}",
	"{{.ContainerConfig.Hostname}}",
	"{{json .ContainerConfig.Entrypoint}}",
	"{{json .ContainerConfig.Cmd}}",
	"{{json .ContainerConfig.Env}}",
}
