require.config({
    baseUrl: '/js/',
    urlArgs: "bust=" + (new Date()).getTime(),
    paths: {
        'jquery': 'externals/jquery/jquery-2.0.3.min',
        'jqueryuiwidget': 'externals/jqueryui/jquery.ui.widget',
        'bootstrap': 'externals/bootstrap/bootstrap.min',
        'facebookAll': 'externals/facebook/all.min',
        'gapi': 'externals/gplus/plusone.min'
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
        },

        "facebookAll": {
            exports: "FB"
        }
    }
});

require([
        'jquery',
        'modules/twitter/twitter_widget',
        'modules/facebook/facebook_widget',
        'gapi'
    ], 
    function ($) {
        $(document).ready(function() {
            // $(".ps-js-twitter").twitterSync();
            // $(".ps-js-facebook").facebookSync();
        });
    }
);