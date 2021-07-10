import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './setting'
import { _fetch } from "./_base"
import { mapUser } from "./user"
import { mapDataToDeployment } from "./deployment"
import { User, Deployment, Approval, HttpNotFoundError } from '../models'

export const mapDataToApproval = (data: any): Approval => {
    let user: User | null = null
    let deployment: Deployment | null = null

    if ("user" in data.edges) {
        user = mapUser(data.edges.user)
    }

    if ("deployment" in data.edges) {
        deployment = mapDataToDeployment(data.edges.deployment)
    }

    return  {
        isApproved: data.is_approved,
        createdAt: new Date(data.created_at),
        updatedAt: new Date(data.updated_at),
        user,
        deployment
    }
}

export const listApprovals = async (id: string, number: number) => {
    const res = await _fetch(`${instance}/api/v1/repos/${id}/deployments/${number}/approvals`, {
        credentials: "same-origin",
        headers,
    })

    if (res.status === StatusCodes.NOT_FOUND) {
        throw new HttpNotFoundError("There is no requested approval.")
    }

    const approvals: Approval[] = await res.json().
        then(data => data.map((d:any): Approval => mapDataToApproval(d)))

    return approvals
}

export const getApproval = async (id: string, number: number) => {
    const res = await _fetch(`${instance}/api/v1/repos/${id}/deployments/${number}/approval`, {
        credentials: "same-origin",
        headers,
    })

    if (res.status === StatusCodes.NOT_FOUND) {
        throw new HttpNotFoundError("There is no requested approval.")
    }

    const approval = await res.json().
        then(data => mapDataToApproval(data))
    return approval
}

export const setApprovalApproved = async (id: string, number: number) => {
    const body = {
        is_approved: true,
    }
    const res = await _fetch(`${instance}/api/v1/repos/${id}/deployments/${number}/approval`, {
        credentials: "same-origin",
        headers,
        method: "PATCH",
        body: JSON.stringify(body),
    })

    if (res.status === StatusCodes.NOT_FOUND) {
        throw new HttpNotFoundError("There is no requested approval.")
    }

    const approval = await res.json().
        then(data => mapDataToApproval(data))
    return approval
}

export const setApprovalDeclined = async (id: string, number: number) => {
    const body = {
        is_approved: false,
    }
    const res = await _fetch(`${instance}/api/v1/repos/${id}/deployments/${number}/approval`, {
        credentials: "same-origin",
        headers,
        method: "PATCH",
        body: JSON.stringify(body),
    })

    if (res.status === StatusCodes.NOT_FOUND) {
        throw new HttpNotFoundError("There is no requested approval.")
    }

    const approval = await res.json().
        then(data => mapDataToApproval(data))
    return approval
}