document.addEventListener('DOMContentLoaded', function() {
  var form = document.getElementById('Message-Form');
  if (form) {
      form.addEventListener('submit', function(e) {
          e.preventDefault();

          // Récupérer les valeurs des champs du formulaire
          var messageContent = document.getElementById('messageContent').value;
          var author = document.getElementById('author').value;

          // Créer un objet avec les données formatées
          var data = {
              message_content: messageContent,
              author_of_the_message: author
          };

          // Effectuer la requête POST avec les données formatées
          fetch('/api/message', {
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