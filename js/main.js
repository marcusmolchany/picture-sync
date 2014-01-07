require.config({
    baseUrl: '/picture-sync/js/',
    urlArgs: "bust=" + (new Date()).getTime(),
    paths: {
        'jquery': 'jquery/jquery-2.0.3.min',
        'jqueryui': 'jqueryui.ui.widget',
        'bootstrap': 'bootstrap/bootstrap.min',
    },

    // Configure the dependencies and exports for older, traditional "browser globals" - 
    // scripts that do not use define() to declare the dependencies and set a module value.
    shim: {  
        "jqueryuiwidget": {
            exports: "$",
            deps: ['jquery']
        }
    }
});

require([
        'jquery',
        'modules/jquerynewmodule',
        'modules/jqueryoldmodule',
        'modules/myplugin'
    ], 
    function ($, jquerynewmodule, jqueryoldmodule) {
        $(document).ready(function() {

            // Demonstrate multiple versions of same dependency (jquery 1.8, 1.9)
            $(".bst-js-launchdemo-new").click(function(){
                jquerynewmodule.whatVersionIsThisUsing();
            });
            
            $(".bst-js-launchdemo-old").click(function(){
                jqueryoldmodule.whatVersionIsThisUsing();
            });

            // Demonstrate jqueryui widget (plugin) format with Require.js framework
            $(".bst-js-launchdemo-plugin").myplugin({
                modalClass: "bst-js-my-modal",
                modifyClass: "bst-js-modify-this"
            });
        });
    }
);