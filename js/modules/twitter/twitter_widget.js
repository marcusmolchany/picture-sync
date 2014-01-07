define(['jquery', 'jqueryuiwidget'], function ($) {
    $.widget('ps.twitterSync', {
        options: {
            
        },

        _init: function() {
            console.dir("twitter widget set up");
            this.setProfileName();
        },

        setProfileName: function() {
            var myProfileName = "My profile name!";

            // do some ajax things with facebook
            // get the person's profile name
            // put it into myProfileName

            $(this.element).children(".ps-js-profile-name").text(myProfileName);
        }
    });
});