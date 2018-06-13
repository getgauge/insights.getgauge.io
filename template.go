// Copyright 2015 ThoughtWorks, Inc.

// This file is part of Gauge.

// Gauge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge.  If not, see <http://www.gnu.org/licenses/>.

package main

const htmltemplate = `
{{$maxRows := 15}}
{{define "table"}}
    <div class="table-container">
        <table class="u-full-width">
            <thead>
                <tr>
                    <th>Category</th>
                    <th>Action</th>
                    <th>Label</th>
                    <th>Hits</th>
                    <th>%</th>
                </tr>
            </thead>
			<tbody>
				{{range .}}
					<tr>
						<td>{{.Category}}</td>
						<td>{{.Action}}</td>
						<td>{{.Labels}}</td>
						<td>{{.Hits}}</td>
						<td>{{printf "%.2f" .Percent}}</td>
					</tr>
				{{end}}
            </tbody>
        </table>
    </div>
{{end}}
<!DOCTYPE html>
<html>
	<head>
		<title>Gauge - Insights</title>
    	<link href="https://getgauge.io/assets/images/favicons/favicon.ico" rel="shortcut icon" type="image/ico" />
		<link rel="stylesheet" type="text/css" href="https://fonts.googleapis.com/css?family=Raleway:400,300,600">
        <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/skeleton/2.0.4/skeleton.css">
		<style>
			body,
			html {
				text-align: center;
				margin: 0;
				overflow: hidden;
				-webkit-transition: opacity 400ms;
				-moz-transition: opacity 400ms;
				transition: opacity 400ms;
			}

			body,
			.onepage-wrapper,
			html {
				display: block;
				position: static;
				padding: 0;
				width: 100%;
				height: 100%;
			}

			.onepage-wrapper {
				width: 100%;
				height: 100%;
				display: block;
				position: relative;
				padding: 0;
				-webkit-transform-style: preserve-3d;
			}

			.onepage-wrapper .ops-section {
				width: 100%;
				height: 100%;
				position: relative;
			}

			.onepage-pagination {
				position: absolute;
				right: 10px;
				top: 50%;
				z-index: 5;
				list-style: none;
				margin: 0;
				padding: 0;
			}

			.onepage-pagination li {
				padding: 0;
				text-align: center;
			}

			.onepage-pagination li a {
				padding: 10px;
				width: 4px;
				height: 4px;
				display: block;
			}

			.onepage-pagination li a:before {
				content: '';
				position: absolute;
				width: 4px;
				height: 4px;
				background: rgba(0, 0, 0, 0.85);
				border-radius: 10px;
				-webkit-border-radius: 10px;
				-moz-border-radius: 10px;
			}

			.onepage-pagination li a.active:before {
				width: 10px;
				height: 10px;
				background: none;
				border: 1px solid black;
				margin-top: -4px;
				left: 8px;
			}

			.disabled-onepage-scroll,
			.disabled-onepage-scroll .wrapper {
				overflow: auto;
			}

			.disabled-onepage-scroll .onepage-wrapper .ops-section {
				position: relative !important;
				top: auto !important;
			}

			.disabled-onepage-scroll .onepage-wrapper {
				-webkit-transform: none !important;
				-moz-transform: none !important;
				transform: none !important;
				-ms-transform: none !important;
				min-height: 100%;
			}

			.disabled-onepage-scroll .onepage-pagination {
				display: none;
			}

			body.disabled-onepage-scroll,
			.disabled-onepage-scroll .onepage-wrapper,
			html {
				position: inherit;
			}

			.header {
				background: #F5C20F;
				color: #343131;
			}

			.header .heading {
				font-size: 1.8em;
				font-weight: bold;
			}



			.logo {
				height: 4em;
				margin-top: 1em;
			}

			h1 {
				font-size: 2.5em;
				color: #343131;
			}

			h2 {
				font-size: 1.8em;
				margin-bottom: 1rem;
			}

			.table-container {
				display: inline-block;
				overflow: hidden;
				padding-bottom: 1em;
			}

			table {
				border-collapse: collapse;
			}

			td,th {
				padding: 5px 15px;
				text-align: left;
				border-bottom: 1px solid #B58F0B;
				font-size: medium;
			}

			.truncated:after {
				content: '<Truncated. Download CSV for Full data>';
				font-size: 1em;
			}



			tr:nth-child(n+{{$maxRows}}){
				display:none;
			}

			footer {
				margin-left: 22em;
				margin-top: 2em;
				margin-bottom: 1.5em;
			}

			.caption {
				font-size: 2em;
				margin: 0;
			}

			section {
				margin: 0;
				background-color: #F5C20F;
			}

			@media only screen and (min-width: 768px) {
				.mobile {
					display: none;
				}
			}

			@media only screen and (max-width: 768px) {
				.content {
					display: none;
				}
				.mobile {
					margin-top: 3em;
				}
			}
			#map{
				height: 70em;
				padding-top: 6em;
			}

			.reportPeriod {
				font-weight: bold;
			}

			.usages {
				padding: 2rem;
			}

			.download {
				color: #755C07;
				border: 1px solid #755C07;
			}

			.download:hover {
				font-weight: bold;
				color: #362B03;
				border: 1px solid #362B03;
			}

			.gauge-logo {
				background-image: url(data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz48c3ZnIHZlcnNpb249IjEuMSIgaWQ9IkxheWVyXzEiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiIHg9IjBweCIgeT0iMHB4IiB2aWV3Qm94PSIwIDAgMTM2MC42IDQxMSIgc3R5bGU9ImVuYWJsZS1iYWNrZ3JvdW5kOm5ldyAwIDAgMTM2MC42IDQxMTsiIHhtbDpzcGFjZT0icHJlc2VydmUiPjxnPjxwYXRoIGQ9Ik01NDcuOSwzMTAuNWMtNC4yLDExLjMtMTAuMiwyMS0xOCwyOS4xYy03LjgsOC4xLTE3LjEsMTQuNC0yOCwxOC44Yy0xMC45LDQuNC0yMi45LDYuNy0zNiw2LjdjLTIwLjQsMC0zNy45LTUuMS01Mi40LTE1LjJjLTE0LjUtMTAuMS0yNC0yMy45LTI4LjUtNDEuNWgyOS42YzIuOSw2LjQsNi41LDExLjYsMTEsMTUuNWM0LjQsMy45LDkuMiw2LjgsMTQuMiw4LjhjNSwyLDkuOCwzLjMsMTQuNSw0czguNSwxLDExLjcsMWMxMC45LDAsMjAuMS0xLjksMjcuOC01LjhjNy43LTMuOSwxMy45LTguOSwxOC42LTE1YzQuOC02LjEsOC4yLTEzLDEwLjMtMjAuOGMyLjEtNy44LDMuMi0xNS41LDMuMi0yMy4zdi0xOGMtNi4yLDguNC0xNC40LDE1LjYtMjQuNSwyMS41Yy0xMC4xLDUuOS0yMS40LDguOC0zMy44LDguOHMtMjMuOS0yLjMtMzQuMy03cy0xOS41LTExLjEtMjcuMS0xOS4zYy03LjctOC4yLTEzLjctMTcuOC0xOC0yOC42Yy00LjMtMTAuOS02LjUtMjIuNS02LjUtMzVzMi4yLTI0LjEsNi41LTM1LjFjNC4zLTExLDEwLjMtMjAuNSwxOC0yOC42YzcuNy04LjEsMTYuNy0xNC41LDI3LjEtMTkuM2MxMC40LTQuOCwyMS45LTcuMiwzNC4zLTcuMmMxMi43LDAsMjQuMiwzLjEsMzQuNiw5LjNTNTIxLDEyOCw1MjcsMTM2Ljd2LTI3LjNoMjcuM3YxNjMuNUM1NTQuMywyODYuNiw1NTIuMiwyOTkuMiw1NDcuOSwzMTAuNXogTTUyMS4zLDE3MC42Yy0zLjEtNy44LTcuMy0xNC42LTEyLjctMjAuNWMtNS4zLTUuOS0xMS41LTEwLjUtMTguNS0xMy44cy0xNC41LTUtMjIuNS01Yy04LDAtMTUuNSwxLjctMjIuNiw1Yy03LjEsMy4zLTEzLjIsNy45LTE4LjMsMTMuOGMtNS4xLDUuOS05LjIsMTIuNy0xMi4yLDIwLjVjLTMsNy44LTQuNSwxNi4xLTQuNSwyNWMwLDguNywxLjUsMTYuOSw0LjUsMjQuNmMzLDcuOCw3LDE0LjUsMTIuMiwyMC4zYzUuMSw1LjgsMTEuMiwxMC4zLDE4LjMsMTMuN2M3LjEsMy4zLDE0LjcsNSwyMi42LDVjOCwwLDE1LjUtMS43LDIyLjUtNWM3LTMuMywxMy4yLTcuOSwxOC41LTEzLjdjNS4zLTUuOCw5LjUtMTIuNSwxMi43LTIwLjNjMy4xLTcuOCw0LjctMTYsNC43LTI0LjZDNTI2LDE4Ni43LDUyNC40LDE3OC40LDUyMS4zLDE3MC42eiIvPjxwYXRoIGQ9Ik03MTkuNCwyODAuOHYtMjcuNmMtNiw4LjktMTQuMiwxNi40LTI0LjYsMjIuNnMtMjIsOS4zLTM0LjYsOS4zYy0xMi40LDAtMjMuOS0yLjMtMzQuMy03cy0xOS41LTExLjEtMjcuMS0xOS4zYy03LjctOC4yLTEzLjctMTcuOC0xOC0yOC42Yy00LjMtMTAuOS02LjUtMjIuNS02LjUtMzVzMi4yLTI0LjEsNi41LTM1LjFjNC4zLTExLDEwLjMtMjAuNSwxOC0yOC42YzcuNy04LjEsMTYuNy0xNC41LDI3LjEtMTkuM2MxMC40LTQuOCwyMS45LTcuMiwzNC4zLTcuMmMxMi43LDAsMjQuMiwzLjEsMzQuNiw5LjNzMTguNiwxMy43LDI0LjYsMjIuM3YtMjcuM2gyNy4zdjE3MS41SDcxOS40eiBNNzEzLjgsMTcxYy0zLjEtNy44LTcuMy0xNC41LTEyLjctMjAuM2MtNS4zLTUuOC0xMS41LTEwLjMtMTguNS0xMy43cy0xNC41LTUtMjIuNS01Yy04LDAtMTUuNSwxLjctMjIuNiw1Yy03LjEsMy4zLTEzLjIsNy45LTE4LjMsMTMuN2MtNS4xLDUuOC05LjIsMTIuNS0xMi4yLDIwLjNjLTMsNy44LTQuNSwxNi00LjUsMjQuNnMxLjUsMTYuOCw0LjUsMjQuM2MzLDcuNSw3LDE0LjIsMTIuMiwxOS44YzUuMSw1LjcsMTEuMiwxMC4yLDE4LjMsMTMuNWM3LjEsMy4zLDE0LjcsNSwyMi42LDVjOCwwLDE1LjUtMS43LDIyLjUtNWM3LTMuMywxMy4yLTcuOCwxOC41LTEzLjVjNS4zLTUuNyw5LjUtMTIuMywxMi43LTE5LjhjMy4xLTcuNSw0LjctMTUuNyw0LjctMjQuM1M3MTYuOSwxNzguNyw3MTMuOCwxNzF6Ii8+PHBhdGggZD0iTTg5MC42LDI4MC44di0xOC42Yy01LjgsNy4xLTEyLjksMTIuNy0yMS4zLDE2LjZjLTguNCw0LTE4LDYtMjguNiw2Yy0xMC40LDAtMTkuOC0xLjgtMjgtNS4zcy0xNS4yLTguNS0yMC44LTE1Yy01LjctNi40LTkuOS0xNC0xMi44LTIyLjhjLTIuOS04LjgtNC4zLTE4LjUtNC4zLTI5LjFWMTA5LjRoMjh2MTA0LjZjMCwxMy44LDMuNiwyNC44LDEwLjgsMzNzMTcuNywxMi4zLDMxLjUsMTIuM2M2LjIsMCwxMi0xLjEsMTcuNS0zLjNjNS40LTIuMiwxMC4yLTUuMywxNC4yLTkuM2M0LTQsNy4yLTguNyw5LjUtMTQuMmMyLjMtNS40LDMuNS0xMS40LDMuNS0xNy44VjEwOS40aDI4LjN2MTcxLjVIODkwLjZ6Ii8+PHBhdGggZD0iTTExMDYuNywzMTAuNWMtNC4yLDExLjMtMTAuMiwyMS0xOCwyOS4xYy03LjgsOC4xLTE3LjEsMTQuNC0yOCwxOC44Yy0xMC45LDQuNC0yMi45LDYuNy0zNiw2LjdjLTIwLjQsMC0zNy45LTUuMS01Mi40LTE1LjJjLTE0LjUtMTAuMS0yNC0yMy45LTI4LjUtNDEuNWgyOS42YzIuOSw2LjQsNi41LDExLjYsMTEsMTUuNWM0LjQsMy45LDkuMiw2LjgsMTQuMiw4LjhjNSwyLDkuOCwzLjMsMTQuNSw0czguNSwxLDExLjcsMWMxMC45LDAsMjAuMS0xLjksMjcuOC01LjhjNy43LTMuOSwxMy45LTguOSwxOC42LTE1YzQuOC02LjEsOC4yLTEzLDEwLjMtMjAuOGMyLjEtNy44LDMuMi0xNS41LDMuMi0yMy4zdi0xOGMtNi4yLDguNC0xNC40LDE1LjYtMjQuNSwyMS41Yy0xMC4xLDUuOS0yMS40LDguOC0zMy44LDguOGMtMTIuNCwwLTIzLjktMi4zLTM0LjMtN3MtMTkuNS0xMS4xLTI3LjEtMTkuM2MtNy43LTguMi0xMy43LTE3LjgtMTgtMjguNmMtNC4zLTEwLjktNi41LTIyLjUtNi41LTM1czIuMi0yNC4xLDYuNS0zNS4xYzQuMy0xMSwxMC4zLTIwLjUsMTgtMjguNmM3LjctOC4xLDE2LjctMTQuNSwyNy4xLTE5LjNjMTAuNC00LjgsMjEuOS03LjIsMzQuMy03LjJjMTIuNywwLDI0LjIsMy4xLDM0LjYsOS4zczE4LjYsMTMuNywyNC42LDIyLjN2LTI3LjNoMjcuM3YxNjMuNUMxMTEzLDI4Ni42LDExMTAuOSwyOTkuMiwxMTA2LjcsMzEwLjV6IE0xMDgwLjEsMTcwLjZjLTMuMS03LjgtNy4zLTE0LjYtMTIuNy0yMC41Yy01LjMtNS45LTExLjUtMTAuNS0xOC41LTEzLjhzLTE0LjUtNS0yMi41LTVjLTgsMC0xNS41LDEuNy0yMi42LDVjLTcuMSwzLjMtMTMuMiw3LjktMTguMywxMy44Yy01LjEsNS45LTkuMiwxMi43LTEyLjIsMjAuNWMtMyw3LjgtNC41LDE2LjEtNC41LDI1YzAsOC43LDEuNSwxNi45LDQuNSwyNC42YzMsNy44LDcsMTQuNSwxMi4yLDIwLjNjNS4xLDUuOCwxMS4yLDEwLjMsMTguMywxMy43YzcuMSwzLjMsMTQuNyw1LDIyLjYsNWM4LDAsMTUuNS0xLjcsMjIuNS01YzctMy4zLDEzLjItNy45LDE4LjUtMTMuN2M1LjMtNS44LDkuNS0xMi41LDEyLjctMjAuM2MzLjEtNy44LDQuNy0xNiw0LjctMjQuNkMxMDg0LjcsMTg2LjcsMTA4My4yLDE3OC40LDEwODAuMSwxNzAuNnoiLz48cGF0aCBkPSJNMTI3MS45LDI3MGMtMTQsOS45LTMxLDE0LjgtNTAuOSwxNC44Yy0xMi40LDAtMjQtMi4zLTM0LjYtN2MtMTAuNy00LjctMTkuOS0xMS0yNy44LTE5Yy03LjktOC0xNC4xLTE3LjQtMTguNi0yOC4zYy00LjYtMTAuOS02LjgtMjIuNi02LjgtMzUuM2MwLTEyLjQsMi4zLTI0LjEsNi44LTM1LjFzMTAuNy0yMC41LDE4LjUtMjguNmM3LjgtOC4xLDE2LjktMTQuNSwyNy4zLTE5LjNjMTAuNC00LjgsMjEuNi03LjIsMzMuNi03LjJjMTEuMSwwLDIxLjksMi4yLDMyLjMsNi43czE5LjcsMTAuOSwyNy44LDE5LjNzMTQuNSwxOC42LDE5LjMsMzAuNmM0LjgsMTIsNy4yLDI1LjYsNy4yLDQxaC0xNDQuMmMwLjksOC45LDMuMiwxNi44LDYuOCwyMy44czguMiwxMi45LDEzLjcsMTcuOGM1LjQsNC45LDExLjUsOC42LDE4LjEsMTEuMmM2LjcsMi42LDEzLjUsMy44LDIwLjYsMy44YzEyLjIsMCwyMi42LTIuOCwzMS4xLTguM2M4LjUtNS41LDE1LjItMTMuMywxOS44LTIzLjNoMzBDMTI5NS44LDI0NiwxMjg1LjksMjYwLjEsMTI3MS45LDI3MHogTTEyNjYuNCwxNTguM2MtMy43LTUuOC03LjktMTAuNy0xMi44LTE0LjhzLTEwLjMtNy4zLTE2LjItOS41cy0xMS45LTMuMy0xOC4xLTMuM2MtNiwwLTExLjksMS4xLTE3LjgsMy4yYy01LjksMi4xLTExLjMsNS4yLTE2LjMsOS4yYy01LDQtOS40LDguOS0xMy4zLDE0LjdjLTMuOSw1LjgtNi43LDEyLjMtOC41LDE5LjZoMTExLjJDMTI3Mi44LDE3MC40LDEyNzAsMTY0LjEsMTI2Ni40LDE1OC4zeiIvPjwvZz48Zz48cGF0aCBkPSJNMjk3LjYsMzI2SDcwLjljLTE1LjcsMC0yOC4zLTEyLjctMjguMy0yOC4zdi05OS4yYzAtMTUuNywxMi43LTI4LjMsMjguMy0yOC4zczI4LjMsMTIuNywyOC4zLDI4LjN2NzAuOWgxMzBMNTAuOCw5MC45Yy0xMS4xLTExLjEtMTEuMS0yOSwwLTQwLjFjMTEuMS0xMS4xLDI5LTExLjEsNDAuMSwwbDIyNi44LDIyNi44YzguMSw4LjEsMTAuNSwyMC4zLDYuMSwzMC45QzMxOS40LDMxOS4xLDMwOS4xLDMyNiwyOTcuNiwzMjZ6Ii8+PHBhdGggZD0iTTI5Ny42LDE4NC4zYy03LjMsMC0xNC41LTIuOC0yMC04LjNsLTg1LTg1Yy0xMS4xLTExLjEtMTEuMS0yOSwwLTQwLjFjMTEuMS0xMS4xLDI5LTExLjEsNDAuMSwwbDg1LDg1YzExLjEsMTEuMSwxMS4xLDI5LDAsNDAuMUMzMTIuMSwxODEuNSwzMDQuOSwxODQuMywyOTcuNiwxODQuM3oiLz48L2c+PC9zdmc+);
				height: 3em;
				background-repeat: no-repeat;
				left: calc(50% - 5em);
				position: relative;
			}
		</style>
	</head>
	<body>
		<div class="mobile">
			<div class="header">
				<div class="gauge-logo"></div>
				<h2 class="heading">Insights</h2>
			</div>
			<p>This page is not supported in this resolution.</p>
		</div>
		<div class="content">
			<section>
				<div class="header">
					<div class="gauge-logo"></div>
					<h2 class="heading">Insights</h2>
				</div>
				<h3 class="caption">
					<span class="reportPeriod">{{startDate}}</span> to <span class="reportPeriod">{{endDate}}</span>
				</h3>
				<span>Report generated on : {{now}}</span>
				<div id="map"></div>
			</section>
			<section>
				<h2 class="usages" id="console">Command Usage - Console</h2>
				{{template "table" .Console}}
				<div class="actions">
					<a class="download button" href="./console.csv">Download</a>
				</div>
			</section>
			<section>
				<h2 class="usages" id="ci">Command Usage - CI</h2>
				{{template "table" .CI}}
				<div class="actions">
					<a class="button download" href="./ci.csv">Download</a>
				</div>
			</section>
			<section>
				<h2 class="usages" id="api">Command Usage - API</h2>
				{{template "table" .API}}
				<div class="actions">
					<a class="button download" href="./api.csv">Download</a>
				</div>
			</section>
		</div>
		<script src="https://cdn.rawgit.com/peachananr/purejs-onepage-scroll/master/onepagescroll.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/d3/3.5.3/d3.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/topojson/1.6.9/topojson.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/datamaps/0.5.8/datamaps.world.min.js"></script>
		<script type="text/javascript">
			(function () {
				if (window.matchMedia("(max-width: 768px)").matches) {
					return;
				}
				var countries = [
					{{range .CountryWiseUsers}}
						{name: {{.Country.Name}}, latitude: {{.Country.Lat}}, longitude: {{.Country.Long}}, count: {{.UserCount}}, radius: {{.Radius}}},
					{{end}}
				];
				var map = new Datamap({
					scope: 'world',
					element: document.getElementById('map'),
					projection: 'mercator',
					width: 750,
					height: 500,
					fills: {
						defaultFill: '#755C07'
					},
					geographyConfig: {
						borderWidth: 1,
						borderOpacity: 1,
						borderColor: '#755C07',
						popupOnHover: false,
						highlightOnHover: false,
					},
					bubblesConfig: {
						popupTemplate: function (geo, data) {
							return "<div class='hoverinfo'>" + data.name + " : " + data.count + "</div>";
						},
						highlightFillColor: '#DBAD0D',
        				highlightBorderColor: '#B58F0B',
					},					
				});
				map.bubbles(countries);
				onePageScroll(".content", { loop: true });
				var tables= document.getElementsByTagName("table");
				for(i=0; i< tables.length; i++){
					var table = tables[i];
					if (table.getElementsByTagName("tr").length > {{$maxRows}}) {
						table.parentElement.classList.add("truncated");
					}
				};
				})();
		</script>
	</body>
</html>
`
