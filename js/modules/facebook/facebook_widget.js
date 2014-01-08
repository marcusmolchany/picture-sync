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
                console.dir("authResponseChange occurred");
                console.dir(response);

                if (response.status == "connected") {
                    FB.api('/me', function(response) {
                        self.setProfileName(response.name);
                    });
                } else {
                    console.dir("Something went wrong. Check out the response above ^");
                }
            });

            FB.login();
        },

        setProfileName: function(name) {
            $(this.element).children(".ps-js-profile-name").text(name);
        }
    });
});