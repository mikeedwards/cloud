function expandApiTypes(base) {
    return {
        'START': base + "_START",
        'SUCCESS': base + "_SUCCESS",
        'FAILURE': base + "_FAILURE",
    };
}

export const API_EXPEDITION_GEOJSON_GET = expandApiTypes("API_EXPEDITION_GEOJSON_GET");
export const API_PROJECT_EXPEDITIONS_GET = expandApiTypes("API_PROJECT_EXPEDITIONS_GET");
export const API_EXPEDITION_GET = expandApiTypes("API_EXPEDITION_GET");
export const API_PROJECT_GET = expandApiTypes("API_PROJECT_GET");

export function isVerboseActionType(type) {
    return false;
}