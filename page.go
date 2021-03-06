package lcars

import ()

type Page struct {
	Content string
}

func pageBuilder(mymenu Menu, mysettings Settings) string {
	lcarstopcolor := mysettings.TopColor
	lcarsbottomcolor := mysettings.BottomColor
	title := mysettings.Title
	pageReturn := `
	<html lang="en">
	<head>
	  <meta charset="UTF_8">
	  <title>` + title + `</title>
	  <style  type="text/css">
	  `
	pageReturn = pageReturn + lcarscss +
		`
		</style>
	  <style>

	  html, body { background: black }
	  p, h1, h2, h3 {
		margin-top: 1em;
	  }

	  pre {
		  display: inline;
		    white-space: pre-line;
		    word-wrap: break-word;
	  }
	  </style>
	</head>
  <body>
    <div class="lcars-app-container">
	<!-- HEADER================================== -->
	<div id="header" class="lcars-row header">
	  <!-- ELBOW -->
	  <div class="lcars-elbow left-bottom lcars-` + lcarstopcolor + `-bg"></div>

	  <!-- BAR WITH TITLE -->
	  <div class="lcars-bar horizontal lcars-` + lcarstopcolor + `-bg">
	    <div class="lcars-title right">` + title + `</div>
	  </div>

	  <!-- ROUNDED EDGE DECORATED -->
	  <div class="lcars-bar horizontal right-end decorated lcars-` + lcarstopcolor + `-bg"></div>
	</div>

	<!-- SIDE MENU ================== -->

	<div id="left-menu" class="lcars-column start-space lcars-u-1">
	`
	if mymenu.Items != nil {
		pageReturn = pageReturn + makeMenu(mymenu, mysettings)
	}
	pageReturn = pageReturn + `

	  <!-- FILLER -->
	  <div class="lcars-bar lcars-` + lcarsbottomcolor + `-bg lcars-u-1"></div>
	</div>

	<!-- FOOTER ============================ -->

	<div id="footer" class="lcars-row ">
	  <!-- ELBOW -->
	  <div class="lcars-elbow left-top lcars-` + lcarsbottomcolor + `-bg"></div>
	  <!-- BAR -->
	  <div class="lcars-bar horizontal lcars-` + lcarsbottomcolor + `-bg both-divider bottom"></div>
	  <!-- ROUNDED EDGE -->
	  <div class="lcars-bar horizontal lcars-` + lcarsbottomcolor + `-bg right-end left-divider bottom"></div>
	</div>

	<!-- MAIN CONTAINER -->
	<div id="container">
	  <!-- COLUMN LAYOUT -->
	  <div class="lcars-column lcars-u-5">

	    {{ .Content }}

	  </div>
	</div>
  </div>

  </body>
  </html>
	`
	return pageReturn
}
