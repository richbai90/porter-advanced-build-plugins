package build

import (
	plugin "github.com/hashicorp/go-plugin"
	"github.com/richbai90/porter-advanced-build-plugins/pkg/config"
)



func NewPlugin(cfg config.Config) plugin.Plugin {
	return &Plugin{
		Impl: NewBuilder(cfg),
	}
}


/*
func (b *Plugin) BuildInvocationImage(manifest *manifest.Manifest) error {
	fmt.Fprintf(b.Out, "\nStarting Invocation Image Build (%s) =======> \n", manifest.Image)
	buildOptions := types.ImageBuildOptions{
		SuppressOutput: false,
		PullParent:     false,
		Remove:         true,
		Tags:           []string{manifest.Image},
		Dockerfile:     filepath.ToSlash(b.DOCKER_FILE),
		BuildArgs: map[string]*string{
			"BUNDLE_DIR": &b.BUNDLE_DIR,
		},
	}

	excludes, err := clibuild.ReadDockerignore(b.Getwd())
	if err != nil {
		return err
	}
	excludes = clibuild.TrimBuildFilesFromExcludes(excludes, buildOptions.Dockerfile, false)

	tar, err := archive.TarWithOptions(b.Getwd(), &archive.TarOptions{ExcludePatterns: excludes})
	if err != nil {
		return err
	}

	cli, err := command.NewDockerCli()
	if err != nil {
		return errors.Wrap(err, "could not create new docker client")
	}
	if err := cli.Initialize(cliflags.NewClientOptions()); err != nil {
		return err
	}

	response, err := cli.Client().ImageBuild(context.Background(), tar, buildOptions)
	if err != nil {
		return err
	}

	dockerOutput := ioutil.Discard
	if b.IsVerbose() {
		dockerOutput = b.Out
	}

	termFd, _ := term.GetFdInfo(dockerOutput)
	// Setting this to false here because Moby os.Exit(1) all over the place and this fails on WSL (only)
	// when Term is true.
	isTerm := false
	err = jsonmessage.DisplayJSONMessagesStream(response.Body, dockerOutput, termFd, isTerm, nil)
	if err != nil {
		return errors.Wrap(err, "failed to stream docker build output")
	}
	return nil
}
*/