window.addEventListener('DOMContentLoaded', (event) => {
  fetch('/api/message')
      .then(response => {
          if (!response.ok) {
              throw new Error('Erreur lors de la récupération des données.');
          }
          return response.json();
      })
      .then(data => {
          const container = document.getElementById('topic1-messages');
          const table = document.createElement('table');

          console.log(data);
          const thead = table.createTHead(); // Créez une section <thead> pour les en-têtes de colonne
          const headerRow = thead.insertRow(); // Créez une nouvelle ligne pour les en-têtes de colonne

          // Ajoutez les titres des colonnes
          const headerCell1 = headerRow.insertCell();
          headerCell1.innerHTML = "<b class='fs-2'>#</b>";

          const headerCell2 = headerRow.insertCell();
          headerCell2.innerHTML = "<b class='fs-2'>author</b>";

          const headerCell3 = headerRow.insertCell();
          headerCell3.innerHTML = "<b class='fs-2'>Message Content</b>";

          const headerCell4 = headerRow.insertCell();
          headerCell4.innerHTML = "<b class='fs-2'>Creation Date</b>";

          const headerCell5 = headerRow.insertCell();
          headerCell5.innerHTML = "<b class='fs-2'>vote</b>";

          const tbody = table.createTBody(); // Créez une section <tbody> pour les lignes de données
          data.forEach(message => {
              console.log(message);
              const options = {
                  year: 'numeric',
                  month: 'long',
                  day: 'numeric',
                  hour: 'numeric',
                  minute: 'numeric',
                  second: 'numeric'
              };
              const createdDate = new Date(message.CreatedAt);
              const formattedDate = createdDate.toLocaleString('fr-FR', options); // Formatez la date selon les paramètres régionaux de l'utilisateur

              const row = tbody.insertRow();
              const cell1 = row.insertCell();
              const cell2 = row.insertCell();
              const cell3 = row.insertCell();
              const cell4 = row.insertCell();
              const cell5 = row.insertCell();
              const cell6 = row.insertCell();

              cell1.textContent = message.id;
              cell2.textContent = message.author_of_the_message;
              cell3.textContent = message.message_content;
              cell4.textContent = formattedDate;
              const upvoteButton = document.createElement('button');
              upvoteButton.textContent = 'Upvote';
              upvoteButton.addEventListener('click', function() {
                  // Logique de l'upvote
                  console.log('Upvoted!');
              });

              const downvoteButton = document.createElement('button');
              downvoteButton.textContent = 'Downvote';
              downvoteButton.addEventListener('click', function() {
                  // Logique du downvote
                  console.log('Downvoted!');
              });

              cell5.appendChild(upvoteButton);
              cell5.appendChild(downvoteButton);
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