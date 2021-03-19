var tbody = document.querySelector("#problem_list_wrapper > table > tbody");
var rows = tbody.querySelectorAll("tr");
for(let i = 0; i < rows.length; i++) {
   let row = rows[i];
   let diff = parseFloat(row.children[8].innerText);
    if(diff < 1.5) {
        rows[i].parentNode.removeChild(rows[i]);
    }
   
  }
for(let i = 0; i < rows.length; i++) {
   let row = rows[i];
   let diff = parseFloat(row.children[8].innerText);
    if(diff < 1.5) {
        rows[i].parentNode.removeChild(rows[i]);
    }
   
  }
