<!DOCTYPE html>
<html lang="en">
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">

        <script data-main="js/main" src="/picture-sync/js/externals/require/require-2.1.9.js"></script>

        <link rel="stylesheet" type="text/css" href="/picture-sync/css/app.css">

        <title>Picture Sync</title>
    </head>
    <body>
        <div id="fb-root"></div>
        <div class="container ps-main">
            <nav class="navbar navbar-default ps-header" role="navigation">
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
                        <span class="sr-only">Toggle navigation</span><span class="icon-bar"></span><span class="icon-bar"></span><span class="icon-bar"></span>
                    </button>
                    <a class="navbar-brand" href="/">Sync</a>
                </div>

                <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                    <ul class="nav navbar-nav">
                        <li><a href="/">Boom</a></li>
                        <li><a href="#">Bam</a></li>
                    </ul>
                </div>
            </nav>

            <div class="container">
                <div class="row">
                    <div class="col-md-4 ps-js-twitter">
                        <h3>Twitter</h3>
                        <img class="ps-max-width" src="img/download.png" alt="..." class="img-rounded">
                        <div class="ps-js-profile-name ps-profile-name ps-centered">Your profile name</div>
                        <button type="button" class="ps-js-upload ps-max-width btn btn-primary">Upload and change</button>
                    </div>
                    <div class="col-md-4 ps-js-facebook">
                        <h3>Facebook</h3>
                        <img class="ps-max-width" src="img/download.png" alt="..." class="img-rounded">
                        <div class="ps-js-profile-name ps-profile-name ps-centered">Your profile name</div>
                        <button type="button" class="ps-js-upload ps-max-width btn btn-success">Upload and change</button>
                    </div>
                    <div class="col-md-4 ps-js-instragram">
                        <h3>Instagram?</h3>
                        <img class="ps-max-width" src="img/download.png" alt="..." class="img-rounded">
                        <div class="ps-js-profile-name ps-profile-name ps-centered">Your profile name</div>
                        <button type="button" class="ps-js-upload ps-max-width btn btn-danger">Upload and change</button>
                    </div>
                </div>
            </div>

            <div class="ps-footer-root"></div>
        </div>
        <div class="ps-footer">
            This is the footer
        </div>
    </body>
</html>