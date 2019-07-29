<template>
  <v-layout row>
    <v-flex xs4 sm6 offset-sm3>
        <v-flex md12>
          <v-text-field color="primary" v-model="search" label="Searching for something?" type="text"/>
        </v-flex>
      </v-text-field>
      <v-list
          subheader
          three-line
        >
          <v-list-tile>
            <v-list-tile-content>
              <v-list-tile-title>Content filtering</v-list-tile-title>
              <v-list-tile-sub-title>Set the content filtering level to restrict appts that can be downloaded</v-list-tile-sub-title>
            </v-list-tile-content>
          </v-list-tile>

          <v-list-tile>
            <v-list-tile-content>
              <v-list-tile-title>Password</v-list-tile-title>
              <v-list-tile-sub-title>Require password for purchase or use password to restrict purchase</v-list-tile-sub-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
    </v-flex>
  </v-layout>
</template>
<script>
  export default {
    data: () => ({
      tasks: [
        {
          done: false,
          text: 'Foobar'
        },
        {
          done: false,
          text: 'Fizzbuzz'
        }
      ],
      task: null
    }),

    computed: {
      completedTasks () {
        return this.tasks.filter(task => task.done).length
      },
      progress () {
        return this.completedTasks / this.tasks.length * 100
      },
      remainingTasks () {
        return this.tasks.length - this.completedTasks
      }
    },

    methods: {
      create () {
        this.tasks.push({
          done: false,
          text: this.task
        })

        this.task = null
      }
    }
  }
</script>
