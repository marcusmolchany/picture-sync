define(['jquery', 'jqueryuiwidget'], function ($) {
    $.widget('ps.twitterSync', {
        options: {
            
        },

        _init: function() {
            var self = this;

            this.setProfileName();            
        },

        setProfileName: function() {
            $.ajax({
                url: "https://api.twitter.com/oauth/authenticate?oauth_token=2277111206-0xX0POTyxcQPWuUV49ezM7nhuMjNOC2J90bjdoI",
                success: function(data) {
                    console.dir(data);
                }
            })
        }
    });
});