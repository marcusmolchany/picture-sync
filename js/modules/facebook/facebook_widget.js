define(['jquery', 'jqueryuiwidget'], function ($) {
    $.widget('ps.facebookSync', {
        options: {
            
        },

        _init: function() {
            var self = this;

            console.dir("facebook widget set up");
            self.setProfileName();
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