$(document).ready(function() {
  $("#age_selector").change(function() {
      var value = $("#age_selector").val()
      $.getJSON('/calc/' + value, function(data) {
        var inner = "<p id=phrase>The youngest person you should date is <strong>" + data["Min"] +
        "</strong> years old, and the oldest person that should date you is <strong>" + data["Max"] +
        "</strong> years old.</p>"
        $("#results").html(inner);
        $("#phrase").effect("highlight", 1000);
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
