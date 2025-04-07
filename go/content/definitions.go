package content

import (
	_ "embed"

	"dbut.sh/pkg/writer"
)

//go:embed templates/swarm-gpus.txt
var text1 string

// Index generates the index page content
func Index() string {
	w := writer.New()

	w.Write("")
	w.Write("")
	w.Write(`
 ██████╗ ██████╗ ██████╗ ██╗   ██╗████████╗██████╗ 
██╔═══██╗██╔══██╗██╔══██╗██║   ██║╚══██╔══╝╚════██╗
██║██╗██║██║  ██║██████╔╝██║   ██║   ██║    █████╔╝
██║██║██║██║  ██║██╔══██╗██║   ██║   ██║   ██╔═══╝ 
╚█║████╔╝██████╔╝██████╔╝╚██████╔╝   ██║   ███████╗
 ╚╝╚═══╝ ╚═════╝ ╚═════╝  ╚═════╝    ╚═╝   ╚══════╝
`, writer.WithAlignment(writer.AlignCenter))
	w.Hr()
	w.Write("software engineer | melbourne", writer.WithAlignment(writer.AlignCenter))
	w.Hr()
	w.Write("")
	w.Write("Welcome to the SSH version of my website.")
	w.Write("")
	w.Write("# Blog posts")
	w.Write("")
	w.Write("- [GPUs on a Docker Swarm](/swarm-gpus)")
	w.Write("")
	w.Write("# Links")
	w.Write("")
	w.Write("- [GitHub](https://github.com/dbut2)")
	w.Write("- [LinkedIn](https://linkedin.com/in/dbut2)")
	w.Write("- [Strava](https://strava.com/athletes/dbut2)")
	w.Write("- [Last.fm](https://www.last.fm/user/dbut2)")
	w.Write("")
	w.Write("")

	return w.String()
}

// GpuSwarm generates the GPU Swarm page content
func GpuSwarm() string {
	w := writer.New()

	w.Tr()
	w.Write(`
 ██████╗ ██████╗ ██████╗ ██╗   ██╗████████╗██████╗ 
██╔═══██╗██╔══██╗██╔══██╗██║   ██║╚══██╔══╝╚════██╗
██║██╗██║██║  ██║██████╔╝██║   ██║   ██║    █████╔╝
██║██║██║██║  ██║██╔══██╗██║   ██║   ██║   ██╔═══╝ 
╚█║████╔╝██████╔╝██████╔╝╚██████╔╝   ██║   ███████╗
 ╚╝╚═══╝ ╚═════╝ ╚═════╝  ╚═════╝    ╚═╝   ╚══════╝
`, writer.WithAlignment(writer.AlignCenter))
	w.Tr()
	w.Write("")
	w.Write(text1, writer.WithEnd(writer.Wrap))

	return w.String()
}
