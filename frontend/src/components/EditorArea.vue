<template>
    <div>
        <h1 class="text-center text-white p-3">
            {{ title ? title : 'Editor Area' }}
        </h1>
        <div class="editor-area"></div>
    </div>
</template>

<script>
export default {
    name: 'NodeEditor',
    props: {
        title: String,
    },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.editor-area {
    height: 100vh;
    height: var(--full-height);
    background: orange;
}
</style>
