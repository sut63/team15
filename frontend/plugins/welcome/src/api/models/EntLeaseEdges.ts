/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeFromJSONTyped,
    EntEmployeeToJSON,
    EntRoomdetail,
    EntRoomdetailFromJSON,
    EntRoomdetailFromJSONTyped,
    EntRoomdetailToJSON,
    EntWifi,
    EntWifiFromJSON,
    EntWifiFromJSONTyped,
    EntWifiToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntLeaseEdges
 */
export interface EntLeaseEdges {
    /**
     * 
     * @type {EntEmployee}
     * @memberof EntLeaseEdges
     */
    employee?: EntEmployee;
    /**
     * 
     * @type {EntRoomdetail}
     * @memberof EntLeaseEdges
     */
    roomdetail?: EntRoomdetail;
    /**
     * 
     * @type {EntWifi}
     * @memberof EntLeaseEdges
     */
    wifi?: EntWifi;
}

export function EntLeaseEdgesFromJSON(json: any): EntLeaseEdges {
    return EntLeaseEdgesFromJSONTyped(json, false);
}

export function EntLeaseEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntLeaseEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'employee': !exists(json, 'Employee') ? undefined : EntEmployeeFromJSON(json['Employee']),
        'roomdetail': !exists(json, 'Roomdetail') ? undefined : EntRoomdetailFromJSON(json['Roomdetail']),
        'wifi': !exists(json, 'Wifi') ? undefined : EntWifiFromJSON(json['Wifi']),
    };
}

export function EntLeaseEdgesToJSON(value?: EntLeaseEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'employee': EntEmployeeToJSON(value.employee),
        'roomdetail': EntRoomdetailToJSON(value.roomdetail),
        'wifi': EntWifiToJSON(value.wifi),
    };
}


