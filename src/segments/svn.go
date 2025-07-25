package segments

import (
	"path"
	"strconv"
	"strings"

	"github.com/jandedobbeleer/oh-my-posh/src/regex"
)

// SvnStatus represents part of the status of a Svn repository
type SvnStatus struct {
	ScmStatus
}

func (s *SvnStatus) add(code string) {
	switch code {
	case "?":
		s.Untracked++
	case "C":
		s.Conflicted++
	case "D":
		s.Deleted++
	case "A":
		s.Added++
	case "M":
		s.Modified++
	case "R", "!":
		s.Moved++
	}
}

func (s *SvnStatus) HasConflicts() bool {
	return s.Conflicted > 0
}

const (
	SVNCOMMAND = "svn"
)

type Svn struct {
	Working *SvnStatus
	Branch  string
	scm
	BaseRev int
}

func (s *Svn) Template() string {
	return " \ue0a0{{.Branch}} r{{.BaseRev}} {{.Working.String}} "
}

func (s *Svn) Enabled() bool {
	if !s.shouldDisplay() {
		return false
	}

	s.setSvnStatus()

	return true
}

func (s *Svn) CacheKey() (string, bool) {
	dir, err := s.env.HasParentFilePath(".svn", true)
	if err != nil {
		return "", false
	}

	return dir.Path, true
}

func (s *Svn) shouldDisplay() bool {
	if !s.hasCommand(SVNCOMMAND) {
		return false
	}

	Svndir, err := s.env.HasParentFilePath(".svn", false)
	if err != nil {
		return false
	}

	if Svndir.IsDir {
		s.mainSCMDir = Svndir.Path
		s.scmDir = Svndir.Path
		// convert the worktree file path to a windows one when in a WSL shared folder
		s.repoRootDir = strings.TrimSuffix(s.convertToWindowsPath(Svndir.Path), "/.svn")
		return true
	}

	// handle worktree
	s.scmDir = Svndir.Path
	dirPointer := strings.Trim(s.env.FileContent(Svndir.Path), " \r\n")
	matches := regex.FindNamedRegexMatch(`^Svndir: (?P<dir>.*)$`, dirPointer)
	if matches != nil && matches["dir"] != "" {
		// if we open a worktree file in a WSL shared folder, we have to convert it back
		// to the mounted path
		s.mainSCMDir = s.convertToLinuxPath(matches["dir"])
	}
	return false
}

func (s *Svn) setSvnStatus() {
	s.BaseRev, _ = strconv.Atoi(s.getSvnCommandOutput("info", "--show-item", "revision"))

	branch := s.getSvnCommandOutput("info", "--show-item", "relative-url")
	if len(branch) > 2 {
		s.Branch = branch[2:]
	}

	statusFormats := s.props.GetKeyValueMap(StatusFormats, map[string]string{})
	s.Working = &SvnStatus{ScmStatus: ScmStatus{Formats: statusFormats}}

	displayStatus := s.props.GetBool(FetchStatus, false)
	if !displayStatus {
		return
	}

	changes := s.getSvnCommandOutput("status")
	if changes == "" {
		return
	}
	lines := strings.SplitSeq(changes, "\n")
	for line := range lines {
		if line == "" {
			continue
		}
		// element is the element from someSlice for where we are
		s.Working.add(line[0:1])
	}
}

func (s *Svn) Repo() string {
	// Get the repository name as the last path element of the repository root URL
	repo := s.getSvnCommandOutput("info", "--show-item", "repos-root-url")
	base := path.Base(repo)

	if base == "." {
		return ""
	}

	return base
}

func (s *Svn) getSvnCommandOutput(command string, args ...string) string {
	args = append([]string{command, s.repoRootDir}, args...)
	val, err := s.env.RunCommand(s.command, args...)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(val)
}
