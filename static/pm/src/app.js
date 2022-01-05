const app = Vue.createApp({
    data() {
        return { count: 4 }
    },
    methods: {
        increment() {
            // `this` 指向该组件实例
            this.count++
        }
    }
})

const vm = app.mount('#app')
