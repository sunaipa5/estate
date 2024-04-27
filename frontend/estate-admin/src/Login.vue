<template>
    <div class="col col-sm">
        <form @submit.prevent="loginRequest">
            <h2>Login</h2>
            <input type="text" class="fit" placeholder="username" ref="username">
            <input type="password" class="fit" placeholder="password" ref="password">
            <button class="btn-md btb flr" type="submit">Login</button>
        </form>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    methods: {
        loginRequest() {
            const username = this.$refs.username.value;
            const password = this.$refs.password.value;

            const postData = {
                Username: username,
                Password: password,
            };


            fetch('/login', {
                method: 'POST',
                redirect: 'manual', // Otomatik yönlendirmeyi engelle
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(postData)
            })
                .then(response => {
               
                    console.log(response.status)
                    console.log(typeof(response.status))
                    if (response.status == 404){
                      alert("404 User not found!")
                    }
                    if (response.headers.get('location')){
                        location = response.headers.get('location')
                    }
                  
                   
                    
                })
                .catch(error => {
                    console.error('There was a problem with your fetch operation:', error);
                });





        }
    }
}
</script>