<template>
	<v-toolbar color="secondary">
		<v-avatar :title="tile" :size="avatarSize" color="grey lighten-4">
			<img src="@/assets/islamic_teachings.png" alt="avatar">
		</v-avatar>
		<v-toolbar-title><font color="#004D40">Islamic Teachings</font></v-toolbar-title>
		<v-spacer></v-spacer>
		<v-toolbar-items class="hidden-sm-and-down">
			<v-btn flat><font color="#004D40">Do Good</font></v-btn>
			<v-btn flat><font color="#004D40">Find Masjid</font></v-btn>
		</v-toolbar-items>
	</v-toolbar>
</template>
<script>
export default {
  name: 'NavBar',
  data () {
    return {
      //
    }
  }
}
</script>
