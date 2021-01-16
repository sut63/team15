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
    EntBedtype,
    EntBedtypeFromJSON,
    EntBedtypeFromJSONTyped,
    EntBedtypeToJSON,
    EntCleaningRoom,
    EntCleaningRoomFromJSON,
    EntCleaningRoomFromJSONTyped,
    EntCleaningRoomToJSON,
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeFromJSONTyped,
    EntEmployeeToJSON,
    EntLease,
    EntLeaseFromJSON,
    EntLeaseFromJSONTyped,
    EntLeaseToJSON,
    EntPetrule,
    EntPetruleFromJSON,
    EntPetruleFromJSONTyped,
    EntPetruleToJSON,
    EntPledge,
    EntPledgeFromJSON,
    EntPledgeFromJSONTyped,
    EntPledgeToJSON,
    EntStaytype,
    EntStaytypeFromJSON,
    EntStaytypeFromJSONTyped,
    EntStaytypeToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRoomdetailEdges
 */
export interface EntRoomdetailEdges {
    /**
     * 
     * @type {EntBedtype}
     * @memberof EntRoomdetailEdges
     */
    bedtype?: EntBedtype;
    /**
     * Cleaningrooms holds the value of the cleaningrooms edge.
     * @type {Array<EntCleaningRoom>}
     * @memberof EntRoomdetailEdges
     */
    cleaningrooms?: Array<EntCleaningRoom>;
    /**
     * 
     * @type {EntEmployee}
     * @memberof EntRoomdetailEdges
     */
    employee?: EntEmployee;
    /**
     * 
     * @type {EntPetrule}
     * @memberof EntRoomdetailEdges
     */
    petrule?: EntPetrule;
    /**
     * 
     * @type {EntPledge}
     * @memberof EntRoomdetailEdges
     */
    pledge?: EntPledge;
    /**
     * 
     * @type {EntLease}
     * @memberof EntRoomdetailEdges
     */
    roomdetails?: EntLease;
    /**
     * 
     * @type {EntStaytype}
     * @memberof EntRoomdetailEdges
     */
    staytype?: EntStaytype;
}

export function EntRoomdetailEdgesFromJSON(json: any): EntRoomdetailEdges {
    return EntRoomdetailEdgesFromJSONTyped(json, false);
}

export function EntRoomdetailEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRoomdetailEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bedtype': !exists(json, 'Bedtype') ? undefined : EntBedtypeFromJSON(json['Bedtype']),
        'cleaningrooms': !exists(json, 'Cleaningrooms') ? undefined : ((json['Cleaningrooms'] as Array<any>).map(EntCleaningRoomFromJSON)),
        'employee': !exists(json, 'Employee') ? undefined : EntEmployeeFromJSON(json['Employee']),
        'petrule': !exists(json, 'Petrule') ? undefined : EntPetruleFromJSON(json['Petrule']),
        'pledge': !exists(json, 'Pledge') ? undefined : EntPledgeFromJSON(json['Pledge']),
        'roomdetails': !exists(json, 'Roomdetails') ? undefined : EntLeaseFromJSON(json['Roomdetails']),
        'staytype': !exists(json, 'Staytype') ? undefined : EntStaytypeFromJSON(json['Staytype']),
    };
}

export function EntRoomdetailEdgesToJSON(value?: EntRoomdetailEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bedtype': EntBedtypeToJSON(value.bedtype),
        'cleaningrooms': value.cleaningrooms === undefined ? undefined : ((value.cleaningrooms as Array<any>).map(EntCleaningRoomToJSON)),
        'employee': EntEmployeeToJSON(value.employee),
        'petrule': EntPetruleToJSON(value.petrule),
        'pledge': EntPledgeToJSON(value.pledge),
        'roomdetails': EntLeaseToJSON(value.roomdetails),
        'staytype': EntStaytypeToJSON(value.staytype),
    };
}


