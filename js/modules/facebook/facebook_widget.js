define(['jquery', 'facebookAll', 'jqueryuiwidget'], function ($, FB) {
    $.widget('ps.facebookSync', {
        options: {
            numericUserId: 100007499673243,
            appId: 634051866658660,
            status: true,
            cookie: false,
            xfbml: true
        },

        _init: function() {
            var self = this;

            FB.init({
                appId: self.options.appId,
                status: self.options.status,
                cookie: self.options.cookie,
                xfbml: self.options.xfbml
            });

            FB.Event.subscribe('auth.authResponseChange', function(response) {
                if (response.status == "connected") {
                    self.setupProfile();
                } else {
                    console.dir("Something went wrong with login. Check out the response above ^");
                }
            });

            FB.login();
        },

        setupProfile: function() {
            var self = this;

            self.setProfileName();
            self.setProfilePicture();
        },

        setProfileName: function() {
            var self = this;

            FB.api('/me', function(response) {
                $(self.element).children(".ps-js-profile-name").text(response.name);
            });
        },

        setProfilePicture: function() {
            var self = this;

            FB.api('/me/picture', function(response) {
                $(self.element).children(".ps-js-profile-pic").attr('src', response.data.url);
            });
        }
    });
});