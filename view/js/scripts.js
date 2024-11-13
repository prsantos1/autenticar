<script>
    function enviarDados() {
        
        const username = document.getElementById('usuario').value;
        const password = document.getElementById('senha').value;

       
        const userData = {
            username: username,
            password: password 
        };

       
        const userJson = JSON.stringify(userData);

        
        fetch('http://localhost:808/login', { // URL do servidor Go
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: userJson
        })
        .then(response => response.json())
        .then(data => {
            console.log("Resposta da API", data);

            if (data.success) {
                // Redireciona para a página de login
                window.location.href = "logado.html"; 
            } else {
                alert("Usuário ou senha inválidos");
            }
        })
        .catch(error => {
            console.log("Erro na requisição", error);
            alert("Erro ao fazer login. Tente novamente mais tarde.");
        });
    }
</script>
