import qs from 'qs';
import Client, { IClient } from 'Entities/Client';
import ClientDelete, { IClientDelete } from 'Entities/ClientDelete';
import ClientUpdate, { IClientUpdate } from 'Entities/ClientUpdate';
import Clients, { IClients } from 'Entities/Clients';
import { ClientsFindResponse } from 'Entities/ClientsFindResponse';

// This file was autogenerated. Please do not change.
// All changes will be overwrited on commit.
export default class ClientsApi {
    static async clientsAdd(client: IClient): Promise<number | string[] | Error> {
        const haveError: string[] = [];
        const clientValid = new Client(client);
        haveError.push(...clientValid.validate());
        if (haveError.length > 0) {
            return Promise.resolve(haveError);
        }
        return await fetch(`/clients/add`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(clientValid.serialize()),
        }).then(async (res) => {
            if (res.status === 200) {
                return res.status;
            } else {
                return new Error(String(res.status));
            }
        })
    }

    static async clientsDelete(clientdelete: IClientDelete): Promise<number | string[] | Error> {
        const haveError: string[] = [];
        const clientdeleteValid = new ClientDelete(clientdelete);
        haveError.push(...clientdeleteValid.validate());
        if (haveError.length > 0) {
            return Promise.resolve(haveError);
        }
        return await fetch(`/clients/delete`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(clientdeleteValid.serialize()),
        }).then(async (res) => {
            if (res.status === 200) {
                return res.status;
            } else {
                return new Error(String(res.status));
            }
        })
    }

    static async clientsFind(ip0?: string): Promise<ClientsFindResponse | Error> {
        const queryParams = {
            ip0: ip0,
        }
        return await fetch(`/clients/find?${qs.stringify(queryParams, { arrayFormat: 'comma' })}`, {
            method: 'GET',
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        })
    }

    static async clientsStatus(): Promise<IClients | Error> {
        return await fetch(`/clients`, {
            method: 'GET',
        }).then(async (res) => {
            if (res.status === 200) {
                return res.json();
            } else {
                return new Error(String(res.status));
            }
        })
    }

    static async clientsUpdate(clientupdate: IClientUpdate): Promise<number | string[] | Error> {
        const haveError: string[] = [];
        const clientupdateValid = new ClientUpdate(clientupdate);
        haveError.push(...clientupdateValid.validate());
        if (haveError.length > 0) {
            return Promise.resolve(haveError);
        }
        return await fetch(`/clients/update`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(clientupdateValid.serialize()),
        }).then(async (res) => {
            if (res.status === 200) {
                return res.status;
            } else {
                return new Error(String(res.status));
            }
        })
    }
}