define(['jquery', 'jqueryuiwidget'], function ($) {
    $.widget('ps.twitterSync', {
        options: {
            
        },

        _init: function() {
            
        },

        setProfileName: function() {
            var myProfileName = "My profile name!";

            // do some ajax things with twitter
            // get the person's profile name
            // put it into myProfileName

            $(this.element).children(".ps-js-profile-name").text(myProfileName);
        }
    });
});