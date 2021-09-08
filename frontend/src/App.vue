<template>
    <div id="app" class="container-fluid">
        <div class="row">
            <div class="col-8 bg-primary">
                <h1 class="text-center text-white p-3">Editor Area</h1>
                <div class="editor-area"></div>
            </div>
            <div class="results-area col-4">
                <h1 class="text-center text-white p-3">Results Area</h1>
                <div class="results-container">
                    <div class="code">

                    </div>
                    <div class="actions">
                        <div class="actions-row row">
                                <button type="button" class="btn btn-danger">Discard</button>
                                <button type="button" class="btn btn-dark">Import</button>
                                <button type="button" class="btn btn-primary align-middle">Save</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
  name: 'App',
  components: {
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
.editor-area {
    height: 100vh;
    height: var(--full-height);
    background: orange;
}
.results-area {
    background: #747474;
}
.results-container {
    height: 100vh;
    height: var(--full-height);
    background: white;
}
.code {
    background: #262626;
    height: 80%;
}
.actions {
    background: lightsalmon;
    height: 20%;
    position: relative;
}
.actions-row {
     position: absolute;
     top: 50%;
     left: 50%;
     transform: translate(-50%, -50%);
     background: black;
}
.actions .col button {
}
</style>
