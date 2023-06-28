window.addEventListener('DOMContentLoaded', (event) => {
  fetch('/api/subject')
      .then(response => {
          if (!response.ok) {
              throw new Error('Erreur lors de la récupération des données.');
          }
          return response.json();
      })
      .then(data => {
          const container = document.getElementById('container-subjects');
          const table = document.createElement('table');

          console.log(data);
          const thead = table.createTHead(); // Créez une section <thead> pour les en-têtes de colonne
          const headerRow = thead.insertRow(); // Créez une nouvelle ligne pour les en-têtes de colonne

          // Ajoutez les titres des colonnes
          const headerCell1 = headerRow.insertCell();
          headerCell1.innerHTML = "<b class='fs-3'>#</b>";

          const headerCell2 = headerRow.insertCell();
          headerCell2.innerHTML = "<b class='fs-3'>Subject Name</b>";

          const headerCell3 = headerRow.insertCell();
          headerCell3.innerHTML = "<b class='fs-3'>Description</b>";

          const headerCell4 = headerRow.insertCell();
          headerCell4.innerHTML = "<b class='fs-3'>Creation Date</b>";

          const headerCell5 = headerRow.insertCell();
          headerCell5.innerHTML = "<b class='fs-3'>nb</b>";

          const headerCell6 = headerRow.insertCell();
          headerCell6.innerHTML = "<b class='fs-3'>visit topic</b>";

          const tbody = table.createTBody(); // Créez une section <tbody> pour les lignes de données
          data.forEach(subject => {
              console.log(subject);
              const options = {
                  year: 'numeric',
                  month: 'long',
                  day: 'numeric',
                  hour: 'numeric',
                  minute: 'numeric',
                  second: 'numeric'
              };
              const createdDate = new Date(subject.CreatedAt);
              const formattedDate = createdDate.toLocaleString('fr-FR', options); // Formatez la date selon les paramètres régionaux de l'utilisateur

              const row = tbody.insertRow();
              const cell1 = row.insertCell();
              const cell2 = row.insertCell();
              const cell3 = row.insertCell();
              const cell4 = row.insertCell();
              const cell5 = row.insertCell();
              const cell6 = row.insertCell();

              cell1.textContent = subject.id;
              cell2.textContent = subject.subject_name;
              cell3.textContent = subject.description;
              cell4.textContent = formattedDate;
              cell5.textContent = subject.quantity_of_messages;
              cell6.innerHTML = `<a href="/topics/message.html" class="btn btn-primary">visit topic</a>`;

              // Ajoutez des classes Bootstrap aux éléments du tableau
              table.classList.add('table', 'table-striped', 'table-hover', 'table-dark');
          });
          console.log(container);
          container.appendChild(table);
      })
      .catch(error => {
          console.error(error);
          // Gérer l'erreur d'une manière appropriée, par exemple, afficher un message d'erreur sur la page.
      });
});