$(document).ready(function() {
  $("#age_selector").change(function() {
      var value = $("#age_selector").val()
      $.getJSON('http://localhost:9080/calc/' + value, function(data) {
        var inner = "<p id=phrase>The youngest person you should date is <strong>" + data["Min"] +
        "</strong> years old, and the oldest person that should date you is <strong>" + data["Max"] +
        "</strong> years old.</p>"
        $("#results").html(inner);
        $("#phrase").effect("highlight", 1000);
      });
  });
});
