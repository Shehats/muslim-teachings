import Vue from 'vue'
import './plugins/vuetify'
import App from './App.vue'
import router from './router'
import store from './store'
import Vuetify from 'vuetify'
// Helpers
import colors from 'vuetify/es5/util/colors'

Vue.config.productionTip = false;

Vue.use(Vuetify, {
    theme: {
        primary: colors.teal.darken4,
        secondary: colors.teal.lighten5,
        accent: colors.teal.accent1
    }
})

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
