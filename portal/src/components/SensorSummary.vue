<template>
    <div id="sensor-summary-container" v-if="this.sensor">
        <div id="sensor-title">{{ this.sensor.name }} Statistics</div>
        <p>{{ this.sensor.currentReading ? this.sensor.currentReading.toFixed(1) : "ncR" }}</p>
    </div>
</template>

<script>
export default {
    name: "SensorSummary",
    props: ["sensor"],
    watch: {
        sensor: function() {}
    },
    data: () => {
        return {};
    },
    methods: {}
};
</script>

<style scoped>
#sensor-summary-container {
    float: left;
    width: 32%;
    margin: 10px;
    padding: 10px;
    border: 1px solid rgb(230, 230, 230);
}
#sensor-title {
    margin: 10px 0;
    padding-bottom: 10px;
    border-bottom: 1px solid rgb(200, 200, 200);
}
</style>