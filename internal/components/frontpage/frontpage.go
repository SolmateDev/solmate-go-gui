// Package frontpage contains the page that's displayed on the Hacker News
// window when it's first started.
package frontpage

import (
	_ "embed"

	"github.com/SolmateDev/solmate-go-gui/internal/gtkutil"
)

//go:embed frontpage.css
var css string

func init() { gtkutil.AddCSS(css) }
