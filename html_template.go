//
package main

var HtmlTemplate string = `<!DOCTYPE html>
<html>
  <head>
	  <title>BT Sheet</title>
	</head>
	<body>
	  <header>
		  <h1>Battletech Sheet</h1>
		</header>
		<p>{{.PersonalData.Player}} playing {{.PersonalData.Name}}</p>
		<p>{{.PersonalData.Affiliation}}</p>
	</body>
</html>`
