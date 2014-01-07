require.config({
    baseUrl: '/picture-sync/js/',
    urlArgs: "bust=" + (new Date()).getTime(),
    paths: {
        'jquery': 'jquery/jquery-2.0.3.min',
        'jqueryuiwidget': 'jqueryui/jquery.ui.widget',
        'bootstrap': 'bootstrap/bootstrap.min',
    },

    // Configure the dependencies and exports for older, traditional "browser globals" - 
    // scripts that do not use define() to declare the dependencies and set a module value.
    shim: {  
        "jqueryuiwidget": {
            exports: "$",
            deps: ['jquery']
        },

        "bootstrap": {
            deps: ['jquery']
        }
    }
});

require([
        'jquery',
        'twitter/twitter_widget',
        'facebook/facebook_widget'
    ], 
    function ($, jquerynewmodule, jqueryoldmodule) {
        $(document).ready(function() {
            $(".ps-js-twitter").twitterSync();
            $(".ps-js-facebook").facebookSync();
        });
    }
);