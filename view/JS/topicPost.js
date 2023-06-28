document.addEventListener('DOMContentLoaded', function() {
  var form = document.getElementById('Topic-Form');
  if (form) {
      form.addEventListener('submit', function(e) {
          e.preventDefault();

          // Récupérer les valeurs des champs du formulaire
          var subjectName = document.getElementById('subjectName').value;
          var description = document.getElementById('description').value;

          // Créer un objet avec les données formatées
          var data = {
              subject_name: subjectName,
              description: description
          };

          // Effectuer la requête POST avec les données formatées
          fetch('/api/subject', {
                  method: 'POST',
                  headers: {
                      'Content-Type': 'application/json'
                  },
                  body: JSON.stringify(data)
              })
              .then(function(response) {
                  if (response.ok) {
                      return response.text();
                  } else {
                      throw new Error('Erreur : ' + response.status);
                  }
              })
              .then(function(data) {
                  console.log(data);
              })
              .catch(function(error) {
                  console.log(error);
              });
      });
  }
});