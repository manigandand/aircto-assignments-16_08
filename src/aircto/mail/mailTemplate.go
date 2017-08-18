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
			  				<strong>Issue No: {{ inc $key}}</strong> <br />
			  				<strong>Issue ID: </strong>: {{ $value.ID }} <br />
			  				<strong>Issue Title: </strong>: {{ $value.Title }} <br />
			  				<strong>Issue Description: </strong>: {{ $value.Description }} <br />
			  				<strong>Issue Created By: </strong>: {{ $value.CreatedBy }} <br />
			  				<strong>Issue Status: </strong>: {{ $value.Status }} <br />
			  				<span>---------------------------------------------------</span><br /> <br />
						{{ end }}

			    	</div>
				</div>
		</div>
	</body>
</html>`
