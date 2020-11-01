// AXIOS GLOBALS
axios.defaults.headers.common['X-Auth-Token'] =
  'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c';



// POST REQUEST
function sendMail() {

   console.log("Je suis ici : sendMail() ---> "+document.getElementById('hostname').innerHTML+'/sendmail');

    axios({
           method: 'post',
           url: document.getElementById('hostname').innerHTML+'/sendmail',
           data:  {
                name:     'MyName',
                email:    'email',
                subject:  'Subject',
                comment:  'comment'
             }
            })
         .then(res => console.log(res))
         .catch(err => console.error(err))
         
  }


  // https://stackoverflow.com/questions/46155/how-to-validate-an-email-address-in-javascript
  function validateEmail(email) {
    const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
}


  document.getElementById('submit2').addEventListener('click', sendMail);

