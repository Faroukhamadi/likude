// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Faroukhamadi/likude/ent/comment"
	"github.com/Faroukhamadi/likude/ent/community"
	"github.com/Faroukhamadi/likude/ent/post"
	"github.com/Faroukhamadi/likude/ent/reply"
	"github.com/Faroukhamadi/likude/ent/schema"
	"github.com/Faroukhamadi/likude/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commentMixin := schema.Comment{}.Mixin()
	commentMixinFields0 := commentMixin[0].Fields()
	_ = commentMixinFields0
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescCreatedAt is the schema descriptor for created_at field.
	commentDescCreatedAt := commentMixinFields0[0].Descriptor()
	// comment.DefaultCreatedAt holds the default value on creation for the created_at field.
	comment.DefaultCreatedAt = commentDescCreatedAt.Default.(func() time.Time)
	// commentDescUpdatedAt is the schema descriptor for updated_at field.
	commentDescUpdatedAt := commentMixinFields0[1].Descriptor()
	// comment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	comment.DefaultUpdatedAt = commentDescUpdatedAt.Default.(func() time.Time)
	// comment.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	comment.UpdateDefaultUpdatedAt = commentDescUpdatedAt.UpdateDefault.(func() time.Time)
	communityMixin := schema.Community{}.Mixin()
	communityMixinFields0 := communityMixin[0].Fields()
	_ = communityMixinFields0
	communityFields := schema.Community{}.Fields()
	_ = communityFields
	// communityDescCreatedAt is the schema descriptor for created_at field.
	communityDescCreatedAt := communityMixinFields0[0].Descriptor()
	// community.DefaultCreatedAt holds the default value on creation for the created_at field.
	community.DefaultCreatedAt = communityDescCreatedAt.Default.(func() time.Time)
	// communityDescUpdatedAt is the schema descriptor for updated_at field.
	communityDescUpdatedAt := communityMixinFields0[1].Descriptor()
	// community.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	community.DefaultUpdatedAt = communityDescUpdatedAt.Default.(func() time.Time)
	// community.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	community.UpdateDefaultUpdatedAt = communityDescUpdatedAt.UpdateDefault.(func() time.Time)
	postMixin := schema.Post{}.Mixin()
	postMixinFields0 := postMixin[0].Fields()
	_ = postMixinFields0
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescCreatedAt is the schema descriptor for created_at field.
	postDescCreatedAt := postMixinFields0[0].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the created_at field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	// postDescUpdatedAt is the schema descriptor for updated_at field.
	postDescUpdatedAt := postMixinFields0[1].Descriptor()
	// post.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	post.DefaultUpdatedAt = postDescUpdatedAt.Default.(func() time.Time)
	// post.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	post.UpdateDefaultUpdatedAt = postDescUpdatedAt.UpdateDefault.(func() time.Time)
	replyMixin := schema.Reply{}.Mixin()
	replyMixinFields0 := replyMixin[0].Fields()
	_ = replyMixinFields0
	replyFields := schema.Reply{}.Fields()
	_ = replyFields
	// replyDescCreatedAt is the schema descriptor for created_at field.
	replyDescCreatedAt := replyMixinFields0[0].Descriptor()
	// reply.DefaultCreatedAt holds the default value on creation for the created_at field.
	reply.DefaultCreatedAt = replyDescCreatedAt.Default.(func() time.Time)
	// replyDescUpdatedAt is the schema descriptor for updated_at field.
	replyDescUpdatedAt := replyMixinFields0[1].Descriptor()
	// reply.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	reply.DefaultUpdatedAt = replyDescUpdatedAt.Default.(func() time.Time)
	// reply.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	reply.UpdateDefaultUpdatedAt = replyDescUpdatedAt.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescKarma is the schema descriptor for karma field.
	userDescKarma := userFields[2].Descriptor()
	// user.DefaultKarma holds the default value on creation for the karma field.
	user.DefaultKarma = userDescKarma.Default.(int)
}
