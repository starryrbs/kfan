<script setup>
import {ref, onMounted} from 'vue'
import {Login} from "@/api/account";
import {GetUser, StoreUser} from "@/utils/user";

let userId = ref(-1)
let userName = ref("")
let showLoginInput = ref(false)



onMounted(() => {
  const user = GetUser()
  if (user) {
    const {name, id} = user
    userId.value = id
    userName.value = name
  }
})

const onClickLogin = () => {
  if (!showLoginInput.value) {
    showLoginInput.value = true
  } else {
    Login(userName.value).then(response => {
      userId.value = response.data.id
      showLoginInput.value = false
      StoreUser(userName.value, userId.value)
    })
  }
}
</script>

<template>
  <div>
    <header class="bg-white shadow">
      <div class="px-4 py-6 mx-auto max-w-7xl sm:px-6 lg:px-8">
        <input type="text" v-model="userName" v-if="showLoginInput">

        <a v-if="userId===-1" @click="onClickLogin" style="cursor: pointer">登录</a>
        <a v-else>{{ userName }}</a>
      </div>
    </header>
    <main>
      <router-view/>
    </main>
  </div>
</template>
