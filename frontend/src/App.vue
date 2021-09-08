<template>
    <div id="app" class="container-fluid">
        <div class="row">
            <div class="col-8 bg-primary">
                <EditorArea title="Editor Area" />
            </div>
            <div class="results-area col-4">
                <ResultsArea title="Results Area" />
            </div>
        </div>
    </div>
</template>

<script>
import EditorArea from "./components/EditorArea.vue"
import ResultsArea from "./components/ResultsArea.vue"
export default {
  name: 'App',
  components: {
      EditorArea,
      ResultsArea
  }
}
</script>

<style>
:root {
    --full-height: 88vh;
}
body {
    box-sizing: border-box !important;
    margin: 0;
    display: flex;
    flex-direction: column;
}
#app {
}
.results-area {
    background: #747474;
}
</style>
