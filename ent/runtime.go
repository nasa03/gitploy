// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/hanjunlee/gitploy/ent/approval"
	"github.com/hanjunlee/gitploy/ent/chatcallback"
	"github.com/hanjunlee/gitploy/ent/chatuser"
	"github.com/hanjunlee/gitploy/ent/deployment"
	"github.com/hanjunlee/gitploy/ent/deploymentstatus"
	"github.com/hanjunlee/gitploy/ent/notification"
	"github.com/hanjunlee/gitploy/ent/perm"
	"github.com/hanjunlee/gitploy/ent/repo"
	"github.com/hanjunlee/gitploy/ent/schema"
	"github.com/hanjunlee/gitploy/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	approvalFields := schema.Approval{}.Fields()
	_ = approvalFields
	// approvalDescCreatedAt is the schema descriptor for created_at field.
	approvalDescCreatedAt := approvalFields[1].Descriptor()
	// approval.DefaultCreatedAt holds the default value on creation for the created_at field.
	approval.DefaultCreatedAt = approvalDescCreatedAt.Default.(func() time.Time)
	// approvalDescUpdatedAt is the schema descriptor for updated_at field.
	approvalDescUpdatedAt := approvalFields[2].Descriptor()
	// approval.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	approval.DefaultUpdatedAt = approvalDescUpdatedAt.Default.(func() time.Time)
	// approval.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	approval.UpdateDefaultUpdatedAt = approvalDescUpdatedAt.UpdateDefault.(func() time.Time)
	chatcallbackFields := schema.ChatCallback{}.Fields()
	_ = chatcallbackFields
	// chatcallbackDescIsOpened is the schema descriptor for is_opened field.
	chatcallbackDescIsOpened := chatcallbackFields[2].Descriptor()
	// chatcallback.DefaultIsOpened holds the default value on creation for the is_opened field.
	chatcallback.DefaultIsOpened = chatcallbackDescIsOpened.Default.(bool)
	// chatcallbackDescCreatedAt is the schema descriptor for created_at field.
	chatcallbackDescCreatedAt := chatcallbackFields[3].Descriptor()
	// chatcallback.DefaultCreatedAt holds the default value on creation for the created_at field.
	chatcallback.DefaultCreatedAt = chatcallbackDescCreatedAt.Default.(func() time.Time)
	// chatcallbackDescUpdatedAt is the schema descriptor for updated_at field.
	chatcallbackDescUpdatedAt := chatcallbackFields[4].Descriptor()
	// chatcallback.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	chatcallback.DefaultUpdatedAt = chatcallbackDescUpdatedAt.Default.(func() time.Time)
	// chatcallback.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	chatcallback.UpdateDefaultUpdatedAt = chatcallbackDescUpdatedAt.UpdateDefault.(func() time.Time)
	chatuserFields := schema.ChatUser{}.Fields()
	_ = chatuserFields
	// chatuserDescCreatedAt is the schema descriptor for created_at field.
	chatuserDescCreatedAt := chatuserFields[5].Descriptor()
	// chatuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	chatuser.DefaultCreatedAt = chatuserDescCreatedAt.Default.(func() time.Time)
	// chatuserDescUpdatedAt is the schema descriptor for updated_at field.
	chatuserDescUpdatedAt := chatuserFields[6].Descriptor()
	// chatuser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	chatuser.DefaultUpdatedAt = chatuserDescUpdatedAt.Default.(func() time.Time)
	// chatuser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	chatuser.UpdateDefaultUpdatedAt = chatuserDescUpdatedAt.UpdateDefault.(func() time.Time)
	deploymentFields := schema.Deployment{}.Fields()
	_ = deploymentFields
	// deploymentDescIsRollback is the schema descriptor for is_rollback field.
	deploymentDescIsRollback := deploymentFields[7].Descriptor()
	// deployment.DefaultIsRollback holds the default value on creation for the is_rollback field.
	deployment.DefaultIsRollback = deploymentDescIsRollback.Default.(bool)
	// deploymentDescIsApprovalEnabled is the schema descriptor for is_approval_enabled field.
	deploymentDescIsApprovalEnabled := deploymentFields[8].Descriptor()
	// deployment.DefaultIsApprovalEnabled holds the default value on creation for the is_approval_enabled field.
	deployment.DefaultIsApprovalEnabled = deploymentDescIsApprovalEnabled.Default.(bool)
	// deploymentDescRequiredApprovalCount is the schema descriptor for required_approval_count field.
	deploymentDescRequiredApprovalCount := deploymentFields[9].Descriptor()
	// deployment.DefaultRequiredApprovalCount holds the default value on creation for the required_approval_count field.
	deployment.DefaultRequiredApprovalCount = deploymentDescRequiredApprovalCount.Default.(int)
	// deploymentDescCreatedAt is the schema descriptor for created_at field.
	deploymentDescCreatedAt := deploymentFields[10].Descriptor()
	// deployment.DefaultCreatedAt holds the default value on creation for the created_at field.
	deployment.DefaultCreatedAt = deploymentDescCreatedAt.Default.(func() time.Time)
	// deploymentDescUpdatedAt is the schema descriptor for updated_at field.
	deploymentDescUpdatedAt := deploymentFields[11].Descriptor()
	// deployment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	deployment.DefaultUpdatedAt = deploymentDescUpdatedAt.Default.(func() time.Time)
	// deployment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	deployment.UpdateDefaultUpdatedAt = deploymentDescUpdatedAt.UpdateDefault.(func() time.Time)
	deploymentstatusFields := schema.DeploymentStatus{}.Fields()
	_ = deploymentstatusFields
	// deploymentstatusDescCreatedAt is the schema descriptor for created_at field.
	deploymentstatusDescCreatedAt := deploymentstatusFields[3].Descriptor()
	// deploymentstatus.DefaultCreatedAt holds the default value on creation for the created_at field.
	deploymentstatus.DefaultCreatedAt = deploymentstatusDescCreatedAt.Default.(func() time.Time)
	// deploymentstatusDescUpdatedAt is the schema descriptor for updated_at field.
	deploymentstatusDescUpdatedAt := deploymentstatusFields[4].Descriptor()
	// deploymentstatus.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	deploymentstatus.DefaultUpdatedAt = deploymentstatusDescUpdatedAt.Default.(func() time.Time)
	// deploymentstatus.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	deploymentstatus.UpdateDefaultUpdatedAt = deploymentstatusDescUpdatedAt.UpdateDefault.(func() time.Time)
	notificationFields := schema.Notification{}.Fields()
	_ = notificationFields
	// notificationDescNotified is the schema descriptor for notified field.
	notificationDescNotified := notificationFields[11].Descriptor()
	// notification.DefaultNotified holds the default value on creation for the notified field.
	notification.DefaultNotified = notificationDescNotified.Default.(bool)
	// notificationDescChecked is the schema descriptor for checked field.
	notificationDescChecked := notificationFields[12].Descriptor()
	// notification.DefaultChecked holds the default value on creation for the checked field.
	notification.DefaultChecked = notificationDescChecked.Default.(bool)
	// notificationDescCreatedAt is the schema descriptor for created_at field.
	notificationDescCreatedAt := notificationFields[13].Descriptor()
	// notification.DefaultCreatedAt holds the default value on creation for the created_at field.
	notification.DefaultCreatedAt = notificationDescCreatedAt.Default.(func() time.Time)
	// notificationDescUpdatedAt is the schema descriptor for updated_at field.
	notificationDescUpdatedAt := notificationFields[14].Descriptor()
	// notification.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notification.DefaultUpdatedAt = notificationDescUpdatedAt.Default.(func() time.Time)
	// notification.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notification.UpdateDefaultUpdatedAt = notificationDescUpdatedAt.UpdateDefault.(func() time.Time)
	permFields := schema.Perm{}.Fields()
	_ = permFields
	// permDescCreatedAt is the schema descriptor for created_at field.
	permDescCreatedAt := permFields[2].Descriptor()
	// perm.DefaultCreatedAt holds the default value on creation for the created_at field.
	perm.DefaultCreatedAt = permDescCreatedAt.Default.(func() time.Time)
	// permDescUpdatedAt is the schema descriptor for updated_at field.
	permDescUpdatedAt := permFields[3].Descriptor()
	// perm.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	perm.DefaultUpdatedAt = permDescUpdatedAt.Default.(func() time.Time)
	// perm.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	perm.UpdateDefaultUpdatedAt = permDescUpdatedAt.UpdateDefault.(func() time.Time)
	repoFields := schema.Repo{}.Fields()
	_ = repoFields
	// repoDescConfigPath is the schema descriptor for config_path field.
	repoDescConfigPath := repoFields[4].Descriptor()
	// repo.DefaultConfigPath holds the default value on creation for the config_path field.
	repo.DefaultConfigPath = repoDescConfigPath.Default.(string)
	// repoDescActive is the schema descriptor for active field.
	repoDescActive := repoFields[5].Descriptor()
	// repo.DefaultActive holds the default value on creation for the active field.
	repo.DefaultActive = repoDescActive.Default.(bool)
	// repoDescCreatedAt is the schema descriptor for created_at field.
	repoDescCreatedAt := repoFields[8].Descriptor()
	// repo.DefaultCreatedAt holds the default value on creation for the created_at field.
	repo.DefaultCreatedAt = repoDescCreatedAt.Default.(func() time.Time)
	// repoDescUpdatedAt is the schema descriptor for updated_at field.
	repoDescUpdatedAt := repoFields[9].Descriptor()
	// repo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	repo.DefaultUpdatedAt = repoDescUpdatedAt.Default.(func() time.Time)
	// repo.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	repo.UpdateDefaultUpdatedAt = repoDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAdmin is the schema descriptor for admin field.
	userDescAdmin := userFields[3].Descriptor()
	// user.DefaultAdmin holds the default value on creation for the admin field.
	user.DefaultAdmin = userDescAdmin.Default.(bool)
	// userDescHash is the schema descriptor for hash field.
	userDescHash := userFields[7].Descriptor()
	// user.DefaultHash holds the default value on creation for the hash field.
	user.DefaultHash = userDescHash.Default.(func() string)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[9].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[10].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
