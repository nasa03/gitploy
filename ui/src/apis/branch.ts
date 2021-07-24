import { StatusCodes } from 'http-status-codes'

import { instance, headers } from './setting'
import { _fetch } from "./_base"
import { Branch, HttpNotFoundError } from '../models'

interface BranchData {
    name: string
    commit_sha: string
}

const mapDataToBranch = (data: BranchData): Branch => {
    return {
        name: data.name,
        commitSha: data.commit_sha
    }
}

export const listBranches = async (repoId: string, page = 1, perPage = 30): Promise<Branch[]> => {
    const branches: Branch[] = await _fetch(`${instance}/api/v1/repos/${repoId}/branches?page=${page}&per_page=${perPage}`, {
        headers,
        credentials: "same-origin",
    })
        .then(response => response.json())
        .then(branches => branches.map((b: BranchData) => mapDataToBranch(b)))
    
    return branches
}

export const getBranch = async (repoId: string, name: string): Promise<Branch> => {
    const response = await _fetch(`${instance}/api/v1/repos/${repoId}/branches/${name}`, {
        headers,
        credentials: "same-origin",
    })
    if (response.status === StatusCodes.NOT_FOUND) {
        const message = await response.json().then(data => data.message)
        throw new HttpNotFoundError(message)
    }

    const branch:Branch = await response
        .json()
        .then((b: BranchData) => mapDataToBranch(b))
    
    return branch
}