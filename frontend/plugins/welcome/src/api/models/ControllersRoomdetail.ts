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
    bed?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersRoomdetail
     */
    bedtype?: number;
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
    petrule?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersRoomdetail
     */
    phone?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersRoomdetail
     */
    pledge?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersRoomdetail
     */
    roomnumber?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersRoomdetail
     */
    roomprice?: number;
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
    sleep?: number;
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
        
        'bed': !exists(json, 'bed') ? undefined : json['bed'],
        'bedtype': !exists(json, 'bedtype') ? undefined : json['bedtype'],
        'employee': !exists(json, 'employee') ? undefined : json['employee'],
        'petrule': !exists(json, 'petrule') ? undefined : json['petrule'],
        'phone': !exists(json, 'phone') ? undefined : json['phone'],
        'pledge': !exists(json, 'pledge') ? undefined : json['pledge'],
        'roomnumber': !exists(json, 'roomnumber') ? undefined : json['roomnumber'],
        'roomprice': !exists(json, 'roomprice') ? undefined : json['roomprice'],
        'roomtypename': !exists(json, 'roomtypename') ? undefined : json['roomtypename'],
        'sleep': !exists(json, 'sleep') ? undefined : json['sleep'],
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
        
        'bed': value.bed,
        'bedtype': value.bedtype,
        'employee': value.employee,
        'petrule': value.petrule,
        'phone': value.phone,
        'pledge': value.pledge,
        'roomnumber': value.roomnumber,
        'roomprice': value.roomprice,
        'roomtypename': value.roomtypename,
        'sleep': value.sleep,
        'staytype': value.staytype,
    };
}


