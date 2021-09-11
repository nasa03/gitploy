import { createSlice, PayloadAction, createAsyncThunk } from '@reduxjs/toolkit'
import { message } from "antd"

import { 
    User, 
    Repo, 
    Deployment,
    Branch, 
    Commit, 
    Tag, 
    Config,
    Env,
    Status,
    DeploymentType, 
    RequestStatus, 
    HttpNotFoundError, 
    HttpForbiddenError,
    HttpConflictError,
    HttpUnprocessableEntityError
} from '../models'
import { 
    searchRepo, 
    listPerms,
    getConfig, 
    listBranches, 
    getBranch, 
    listCommits, 
    getCommit, 
    listStatuses, 
    listTags, 
    getTag, 
    createDeployment,
    createApproval,
} from '../apis'

// fetch all at the first page.
const firstPage = 1
const perPage = 100

interface RepoDeployState {
    display: boolean
    repo?: Repo
    config?: Config 
    env?: Env
    envs: Env[]
    type?: DeploymentType 
    branch?: Branch 
    branchStatuses: Status[]
    branches: Branch[]
    commit?: Commit 
    commitStatuses: Status[]
    commits: Commit[]
    tag?: Tag 
    tagStatuses: Status[]
    tags: Tag[]
    /**
     * Approval selecter.
     * approvalEnabled - The approvers field is displayed when it is enabled.
     * approvers - selected approvers from candidates.
    */
    approvers: User[]
    candidates: User[]
    deploying: RequestStatus
    deployId: string
}

const initialState: RepoDeployState = {
    display: false,
    envs: [],
    branchStatuses: [],
    branches: [],
    commitStatuses: [],
    commits: [],
    tagStatuses: [],
    tags: [],
    approvers: [],
    candidates: [],
    deploying: RequestStatus.Idle,
    deployId: "",
}

export const init = createAsyncThunk<Repo, {namespace: string, name: string}, { state: {repoDeploy: RepoDeployState} }>(
    'repoDeploy/init', 
    async (params) => {
        const repo = await searchRepo(params.namespace, params.name)
        return repo
    },
)

export const fetchConfig = createAsyncThunk<Config, void, { state: {repoDeploy: RepoDeployState} }>(
    "repoDeploy/fetchConfig", 
    async (_, { getState, rejectWithValue } ) => {
        const { repo } = getState().repoDeploy
        if (!repo) throw new Error("The repo is not set.")

        try {
            const config = await getConfig(repo.id)
            return config
        } catch (e) {
            return rejectWithValue(e)
        }
    },
)

export const fetchBranches = createAsyncThunk<Branch[], void, { state: {repoDeploy: RepoDeployState }}>(
    "repoDeploy/fetchBranches",
    async (_, { getState }) => {
        const { repo } = getState().repoDeploy
        if (!repo) throw new Error("The repo is not set.")

        const branches = await listBranches(repo.id, firstPage, perPage)
        return branches
    }
)

export const checkBranch = createAsyncThunk<Status[], void, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/checkBranch",
    async (_, { getState }) => {
        const { repo, branch } = getState().repoDeploy
        if (!repo || !branch) throw new Error("The repo and branch are not set.") 

        const { statuses } = await listStatuses(repo.id, branch.commitSha)
        return statuses
    }
)

export const addBranchManually = createAsyncThunk<Branch, string, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/addBranchManually",
    async (name: string, { getState, rejectWithValue }) => {
        const { repo } = getState().repoDeploy
        if (!repo) throw new Error("The repo is not set.")

        try {
            const branch = await getBranch(repo.id, name)
            return branch
        } catch(e) {
            if (e instanceof HttpNotFoundError) {
                message.error("The ref is not found. Check the ref is corrent.")
            } else {
                message.error("It has failed to add the ref.")
            }
            return rejectWithValue(e)
        }
    }
)

export const fetchCommits = createAsyncThunk<Commit[], void, { state: {repoDeploy: RepoDeployState }}>(
    "repoDeploy/fetchCommits",
    async (_, { getState }) => {
        const { repo, branch } = getState().repoDeploy
        if (!repo) throw new Error("The repo is not set.")

        const name = (branch)? branch.name : ""
        const commits = await listCommits(repo.id, name, firstPage, perPage)
        return commits
    }
)

export const checkCommit = createAsyncThunk<Status[], void, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/checkCommit",
    async (_, { getState }) => {
        const { repo, commit } = getState().repoDeploy
        if (!repo || !commit) throw new Error("The repo and commit are not set.") 

        const { statuses } = await listStatuses(repo.id, commit.sha)
        return statuses
    }
)

export const addCommitManually = createAsyncThunk<Commit, string, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/addCommitManually",
    async (sha: string, { getState, rejectWithValue }) => {
        const { repo } = getState().repoDeploy
        if (!repo) throw new Error("The repo is not set.")

        try {
            const commit = await getCommit(repo.id, sha)
            return commit
        } catch(e) {
            if (e instanceof HttpNotFoundError) {
                message.error("The ref is not found. Check the ref is corrent.")
            } else {
                message.error("It has failed to add the ref.")
            }
            return rejectWithValue(e)
        }
    }
)

export const fetchTags = createAsyncThunk<Tag[], void, { state: {repoDeploy: RepoDeployState }}>(
    "repoDeploy/fetchTags",
    async (_, { getState }) => {
        const { repo } = getState().repoDeploy
        if (!repo) throw new Error("The repo is not set.")

        const tags = await listTags(repo.id, firstPage, perPage)
        return tags
    }
)

export const checkTag = createAsyncThunk<Status[], void, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/checkTag",
    async (_, { getState }) => {
        const { repo, tag } = getState().repoDeploy
        if (!repo || !tag) throw new Error("The repo and tag are not set.") 

        const { statuses } = await listStatuses(repo.id, tag.commitSha)
        return statuses
    }
)

export const addTagManually = createAsyncThunk<Tag, string, { state: {repoDeploy: RepoDeployState}}>(
    "repoDeploy/addTagManually",
    async (name: string, { getState, rejectWithValue }) => {
        const { repo } = getState().repoDeploy
        if (!repo) throw new Error("The repo is not set.")

        try {
            const tag = await getTag(repo.id, name)
            return tag
        } catch(e) {
            if (e instanceof HttpNotFoundError) {
                message.error("The ref is not found. Check the ref is corrent.")
            } else {
                message.error("It has failed to add the ref.")
            }
            return rejectWithValue(e)
        }
    }
)

export const searchCandidates = createAsyncThunk<User[], string, { state: {repoDeploy: RepoDeployState }}>(
    "repoDeploy/searchCandidates",
    async (q, { getState, rejectWithValue }) => {
        const { repo } = getState().repoDeploy
        if (!repo) {
            throw new Error("The repo is not set.")
        }

        try {
            const perms = await listPerms(repo, q)
            const candidates = perms.map((p) => {
                return p.user
            })
            return candidates
        } catch(e) {
            return rejectWithValue(e)
        }
    }
)

export const deploy = createAsyncThunk<void, void, { state: {repoDeploy: RepoDeployState}}> (
    "repoDeploy/deploy",
    async (_ , { getState, rejectWithValue, requestId }) => {
        const { repo, env, type, branch, commit, tag, approvers, deploying, deployId } = getState().repoDeploy
        if (!repo) {
            throw new Error("The repo is not set.")
        }
        if (!env) {
            throw new Error("The env is not set.")
        }
        if (deploying !== RequestStatus.Pending || requestId !== deployId ) {
            return
        }

        try {
            let deployment: Deployment
            if (type === DeploymentType.Commit && commit) {
                deployment = await createDeployment(repo.id, type, commit.sha, env.name)
            } else if (type === DeploymentType.Branch && branch) {
                deployment = await createDeployment(repo.id, type, branch.name, env.name)
            } else if (type === DeploymentType.Tag && tag) {
                deployment = await createDeployment(repo.id, type, tag.name, env.name)
            } else {
                throw new Error("The type should be one of them: commit, branch, and tag.")
            }

            const msg = <span>
                It starts to deploy. <a href={`/${repo.namespace}/${repo.name}/deployments/${deployment.number}`}>#{deployment.number}</a>
            </span>
            if (!env.approval?.enabled) {
                message.success(msg, 3)
                return
            }

            approvers.forEach(async (approver) => {
                await createApproval(repo, deployment, approver)
            })
            message.success(msg, 3)
        } catch(e) {
            if (e instanceof HttpForbiddenError) {
                message.error("Only write permission can deploy.", 3)
            } else if (e instanceof HttpUnprocessableEntityError)  {
                message.error(<span>It is unprocesable entity. Discussions <a href="https://github.com/gitploy-io/gitploy/discussions/64">#64</a></span>, 3)
            } else if (e instanceof HttpConflictError) {
                message.error("It has conflicted, please retry it.", 3)
            }
            return rejectWithValue(e)
        }
    }
)

export const repoDeploySlice = createSlice({
    name: "repoDeploy",
    initialState,
    reducers: {
        setDisplay: (state, action: PayloadAction<boolean>) => {
            state.display = action.payload
        },
        setEnv: (state, action: PayloadAction<Env>) => {
            state.env =  action.payload
        },
        setType: (state, action: PayloadAction<DeploymentType>) => {
            state.type = action.payload
        },
        setBranch: (state, action: PayloadAction<Branch>) => {
            state.branch = action.payload
        },
        setCommit: (state, action: PayloadAction<Commit>) => {
            state.commit = action.payload
        },
        setTag: (state, action: PayloadAction<Tag>) => {
            state.tag = action.payload
        },
        addApprover: (state, action: PayloadAction<User>) => {
            const candidate = action.payload

            // Check already exist or not.
            const approver = state.approvers.find(approver => approver.id === candidate.id)
            if (approver !== undefined) {
                return
            }

            state.approvers.push(candidate)
        },
        deleteApprover: (state, action: PayloadAction<User>) => {
            const candidate = action.payload

            const approvers = state.approvers.filter(approver => approver.id !== candidate.id)
            state.approvers = approvers
        },
    },
    extraReducers: builder => {
        builder
            .addCase(init.fulfilled, (state, action) => {
                state.repo = action.payload
            })
            .addCase(fetchConfig.fulfilled, (state, action) => {
                const config = action.payload
                state.envs = config.envs.map(e => e)
                state.config = config
            })
            .addCase(fetchBranches.fulfilled, (state, action) => {
                state.branches = action.payload
            })
            .addCase(checkBranch.fulfilled, (state, action) => {
                state.branchStatuses = action.payload
            })
            .addCase(addBranchManually.fulfilled, (state, action) => {
                state.branches.unshift(action.payload)
            })
            .addCase(fetchCommits.fulfilled, (state, action) => {
                state.commits = action.payload
            })
            .addCase(checkCommit.fulfilled, (state, action) => {
                state.commitStatuses = action.payload
            })
            .addCase(addCommitManually.fulfilled, (state, action) => {
                state.commits.unshift(action.payload)
            })
            .addCase(fetchTags.fulfilled, (state, action) => {
                state.tags = action.payload
            })
            .addCase(checkTag.fulfilled, (state, action) => {
                state.tagStatuses = action.payload
            })
            .addCase(addTagManually.fulfilled, (state, action) => {
                state.tags.unshift(action.payload)
            })
            .addCase(searchCandidates.pending, (state) => {
                state.candidates = []
            })
            .addCase(searchCandidates.fulfilled, (state, action) => {
                state.candidates = action.payload
            })
            .addCase(deploy.pending, (state, action) => {
                if (state.deploying === RequestStatus.Idle) {
                    state.deploying = RequestStatus.Pending
                    state.deployId = action.meta.requestId
                }
            })
            .addCase(deploy.fulfilled, (state, action) => {
                if (state.deploying === RequestStatus.Pending && state.deployId === action.meta.requestId) {
                    state.deploying = RequestStatus.Idle
                }
            })
            .addCase(deploy.rejected, (state, action) => {
                if (state.deploying === RequestStatus.Pending && state.deployId === action.meta.requestId) {
                    state.deploying = RequestStatus.Idle
                }
            })
    }
})