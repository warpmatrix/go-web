$(document).ready(function() {
    $.ajax({
        url: "/api/getData"
    }).then(function(data) {
       for (let index = 0; index < data.length; index++) {
           const user = data[index];
           var tblRow = $("<tr></tr>")
           tblRow.append($("<td></td>").text(user.fname))
           tblRow.append($("<td></td>").text(user.lname))
           $("tbody").append(tblRow)
       }
    });
});
