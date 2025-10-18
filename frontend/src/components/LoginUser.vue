<script setup lang="ts">
  import InputGroup from 'primevue/inputgroup';
  import InputGroupAddon from 'primevue/inputgroupaddon';
  import InputText from 'primevue/inputtext';
  import FloatLabel from 'primevue/floatlabel';
  import Button from 'primevue/button';
  import {computed, ref} from "vue";
  import axios from "axios";

  // Define the username and password variables to store the user input
  // ref creates a reactive reference for username and password.
  // This means any changes to username.value or password.value will trigger updates wherever these references are used.
  const username = ref<string>('')
  const password = ref<string>('')

  // Here we don’t use “ref”, because it’s not referencing just 1 value, But it must be computed from 2 other values of the username and password.
  // We use the computed function to create a computed property that will be used to disable the login button if either the username or password is empty.
  // computed creates a computed property isLoginDisabled. This property automatically updates whenever its dependencies (in this case, username.value and password.value) change.
  const isLoginDisabled =  computed(() => !username.value || !password.value)

  const handleLogin = async () => {
    const response = await axios.post('http://localhost:8080/v1/login_user', {
      username: username.value,
      password: password.value
    }).then(response => {
      console.log(response.data)
    }).catch(error => {
      console.log(error)
    })

    console.log(response)
  }

</script>

<template>
  <div class="flex flex-column row-gap-5">
    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-user"></i>
      </InputGroupAddon>
      <FloatLabel>
<!--        Use v-model to bind the input fields to the username variable-->
        <InputText id="username" v-model="username" />
        <label for="username">Username</label>
      </FloatLabel>
    </InputGroup>

    <InputGroup>
      <InputGroupAddon>
        <i class="pi pi-lock"></i>
      </InputGroupAddon>
      <FloatLabel>
<!--        Use v-model to bind the input fields to the password variable-->
        <InputText type="password" id="password" v-model="password"/>
        <label for="username">Password</label>
      </FloatLabel>
    </InputGroup>

    <Button label="Login" :disabled="isLoginDisabled" @click="handleLogin"/>
  </div>
</template>

<style scoped>

</style>