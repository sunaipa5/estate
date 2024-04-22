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


            axios.post('/login', postData)
                .then(response => {
                    console.log(response.data.message)
                    alert(response.data.message)
                })
                .catch(error => {
                    console.log("Post Error:", error);
                })
        }
    }
}
</script>