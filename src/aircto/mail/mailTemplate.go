package mail

const Template = `
<!doctype html>
<html>
	<head>
		<title>AirCTO Notification</title>
	</head>
	<body style="background: #fff;">
		<div style="width: 70%;margin: 0 auto;background: #f2f4f7;min-height: 400px;border: 2px solid #247ce2;border-radius: 5px;">

		  	<div style="background: #ff6259;color: #fff;font-size: 25px;text-align: center;padding: 15px;"> AirCTO</div>
				<div style="margin: 20px;">
				  	<p style="font-size: 18px;">Hello {{.Name}}</p>
				  	<p style="font-size: 15px;"> {{.Message}} <br>Issue Details:<br><br></p>
				  	<div>
			  			{{ range $key, $value := .Issue }}

			  				<strong>{{ $key}}</strong>: {{ $value }} <br />

						{{ end }}
						<span>---------------------------------------------------</span><br /> <br />
			    	</div>
				</div>
		</div>
	</body>
</html>`
