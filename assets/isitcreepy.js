$(document).ready(function() {
  $("#age_selector").change(function() {
      var value = $("#age_selector").val()
      $.getJSON('/calc/' + value, function(data) {
        var phrase = "My non creepy dating age range is people from " + data["Min"] +
                     " to " + data["Max"] +
                     " years old. Check yours here:";
        // Tweet things
        var tweet_button = '<iframe allowtransparency="true" frameborder="0" scrolling="no" src="https://platform.twitter.com/widgets/tweet_button.html?url=http://isitcreepy.coredump.io' +
        '&text=' + phrase + '" style="width:100px; height:20px;"></iframe>'
        // FB things
        var fb_button = '<iframe src="//www.facebook.com/plugins/like.php?href=http%3A%2F%2Fisitcreepy.coredump.io&amp;send=false&amp;layout=button_count&amp;width=450&amp;show_faces=false&amp;action=like&amp;colorscheme=light&amp;font&amp;height=21&amp;appId=205829242775807" scrolling="no" frameborder="0" style="border:none; overflow:hidden; width:450px; height:21px;" allowTransparency="true"></iframe>';
        // Other stuff
        var inner = "<p id=phrase>The youngest person you should date is <strong>" + data["Min"] +
        "</strong> years old, and the oldest person that should date you is <strong>" + data["Max"] +
        "</strong> years old.</p>"
        $("#results").html(inner);
        $("#phrase").effect("highlight", 1000);
        $("#tweet_button").html(tweet_button);
        $("#fb_button").html(fb_button);
        $("#gplus_button").html("<g:plusone href='http://isitcreepy.coredump.io'></g:plusone>");
        gapi.plusone.go();
      });
      $.getJSON('/stats/', function(data) {
        $.plot($("#placeholder"), [ { label: "Minimum Age", data: data["Min"]},
                                    { label: "Maximum Age", data: data["Max"]}
                                  ], {
                                    series: {
                                      points: { show: false }
                                    },
                                    xaxis: {
                                      min: 14,
                                      max: 80,
                                      ticks: [ 14, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80] ,
                                      axisLabel: "Your age"
                                    },
                                    yaxis: {
                                      axisLabel: "Non creepy ages",
                                      ticks: [ 14, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80] ,
                                      max: 80
                                    },
                                    grid: { markings: [
                                        { color: "#f00", lineWidth: 2, xaxis: {from: value, to: value}}
                                      ]}
                                  } );
        $("#graphexplain").html("<p>This graph shows that while the minimum age goes up when you get old, you also get a broader age range.</p>")
        $("#graphexplain").effect("highlight", 1000);
        $("#placeholder").effect("highlight", 1000);
      });
  });
});
