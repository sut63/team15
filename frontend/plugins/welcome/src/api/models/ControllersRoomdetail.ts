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
/**
 * 
 * @export
 * @interface ControllersRoomdetail
 */
export interface ControllersRoomdetail {
    /**
     * 
     * @type {number}
     * @memberof ControllersRoomdetail
     */
    employee?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersRoomdetail
     */
    equipment?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersRoomdetail
     */
    facilitie?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersRoomdetail
     */
    nearbyplace?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersRoomdetail
     */
    quantity?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersRoomdetail
     */
    roomprice?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersRoomdetail
     */
    roomtypename?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersRoomdetail
     */
    staytype?: number;
}

export function ControllersRoomdetailFromJSON(json: any): ControllersRoomdetail {
    return ControllersRoomdetailFromJSONTyped(json, false);
}

export function ControllersRoomdetailFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersRoomdetail {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'employee': !exists(json, 'employee') ? undefined : json['employee'],
        'equipment': !exists(json, 'equipment') ? undefined : json['equipment'],
        'facilitie': !exists(json, 'facilitie') ? undefined : json['facilitie'],
        'nearbyplace': !exists(json, 'nearbyplace') ? undefined : json['nearbyplace'],
        'quantity': !exists(json, 'quantity') ? undefined : json['quantity'],
        'roomprice': !exists(json, 'roomprice') ? undefined : json['roomprice'],
        'roomtypename': !exists(json, 'roomtypename') ? undefined : json['roomtypename'],
        'staytype': !exists(json, 'staytype') ? undefined : json['staytype'],
    };
}

export function ControllersRoomdetailToJSON(value?: ControllersRoomdetail | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'employee': value.employee,
        'equipment': value.equipment,
        'facilitie': value.facilitie,
        'nearbyplace': value.nearbyplace,
        'quantity': value.quantity,
        'roomprice': value.roomprice,
        'roomtypename': value.roomtypename,
        'staytype': value.staytype,
    };
}


