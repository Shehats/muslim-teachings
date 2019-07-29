<template>
  <v-app light>
    <NavBar/>
    <Search/>
    <v-content>
    </v-content>
    <v-footer class="pa-3" color="secondary">
      <v-spacer></v-spacer>
      <div>&copy; {{ new Date().getFullYear() }}</div>
    </v-footer>
  </v-app>
</template>

<script>
import NavBar from '@/components/NavBar'
import Search from '@/components/Search'

export default {
  name: 'App',
  components: {
    NavBar,
    Search
  },
  data () {
    return {
      //
    }
  }
}
</script>
