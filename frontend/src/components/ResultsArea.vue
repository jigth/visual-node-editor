<template>
    <div>
        <h1 class="text-center text-white p-3">
            {{ title ? title : 'Results Area' }}
        </h1>
        <div class="results-container">
            <div class="code"></div>
            <div class="actions">
                <div class="actions-row row">
                    <button type="button" class="btn btn-danger">
                        Discard
                    </button>
                    <button type="button" class="btn btn-dark">Import</button>
                    <button type="button" class="btn btn-primary align-middle">
                        Save
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'ResultsArea',
    props: {
        title: String,
    },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
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
