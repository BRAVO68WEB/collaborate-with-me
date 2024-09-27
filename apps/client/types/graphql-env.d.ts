/* eslint-disable */
/* prettier-ignore */

export type introspection_types = {
    'Any': unknown;
    'Boolean': unknown;
    'DateTime': unknown;
    'Float': unknown;
    'ID': unknown;
    'Int': unknown;
    'LoginResponse': { kind: 'OBJECT'; name: 'LoginResponse'; fields: { 'access_token': { name: 'access_token'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'String'; ofType: null; }; } }; 'is_success': { name: 'is_success'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'Boolean'; ofType: null; }; } }; }; };
    'Mutation': { kind: 'OBJECT'; name: 'Mutation'; fields: { 'addExcalidrawObject': { name: 'addExcalidrawObject'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'Workspace'; ofType: null; }; } }; 'addUserToWorkspace': { name: 'addUserToWorkspace'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'Workspace'; ofType: null; }; } }; 'createUser': { name: 'createUser'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'User'; ofType: null; }; } }; 'createWorkspace': { name: 'createWorkspace'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'Workspace'; ofType: null; }; } }; 'deleteWorkspace': { name: 'deleteWorkspace'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'Boolean'; ofType: null; }; } }; 'disableUser': { name: 'disableUser'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'Boolean'; ofType: null; }; } }; 'login': { name: 'login'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'LoginResponse'; ofType: null; }; } }; 'removeExcalidrawObject': { name: 'removeExcalidrawObject'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'Workspace'; ofType: null; }; } }; 'removeUserFromWorkspace': { name: 'removeUserFromWorkspace'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'Workspace'; ofType: null; }; } }; 'singleUpload': { name: 'singleUpload'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'UploadResponse'; ofType: null; }; } }; 'updateUser': { name: 'updateUser'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'User'; ofType: null; }; } }; 'updateWorkspace': { name: 'updateWorkspace'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'Workspace'; ofType: null; }; } }; }; };
    'NewUser': { kind: 'INPUT_OBJECT'; name: 'NewUser'; isOneOf: false; inputFields: [{ name: 'username'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'String'; ofType: null; }; }; defaultValue: null }, { name: 'email'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'String'; ofType: null; }; }; defaultValue: null }, { name: 'password'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'String'; ofType: null; }; }; defaultValue: null }]; };
    'NewWorkspace': { kind: 'INPUT_OBJECT'; name: 'NewWorkspace'; isOneOf: false; inputFields: [{ name: 'name'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'String'; ofType: null; }; }; defaultValue: null }, { name: 'is_public'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'Boolean'; ofType: null; }; }; defaultValue: null }, { name: 'user_id'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'ObjectID'; ofType: null; }; }; defaultValue: null }]; };
    'ObjectID': unknown;
    'Query': { kind: 'OBJECT'; name: 'Query'; fields: { 'me': { name: 'me'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'User'; ofType: null; }; } }; 'user': { name: 'user'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'User'; ofType: null; }; } }; 'users': { name: 'users'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'LIST'; name: never; ofType: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'User'; ofType: null; }; }; }; } }; 'workspace': { name: 'workspace'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'Workspace'; ofType: null; }; } }; 'workspaces': { name: 'workspaces'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'LIST'; name: never; ofType: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'Workspace'; ofType: null; }; }; }; } }; }; };
    'String': unknown;
    'Subscription': { kind: 'OBJECT'; name: 'Subscription'; fields: { 'liveUserUpdates': { name: 'liveUserUpdates'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'User'; ofType: null; }; } }; 'liveWorkspaceCollaborators': { name: 'liveWorkspaceCollaborators'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'LIST'; name: never; ofType: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'User'; ofType: null; }; }; }; } }; 'liveWorkspaceUpdates': { name: 'liveWorkspaceUpdates'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'Any'; ofType: null; }; } }; }; };
    'UpdateUser': { kind: 'INPUT_OBJECT'; name: 'UpdateUser'; isOneOf: false; inputFields: [{ name: 'username'; type: { kind: 'SCALAR'; name: 'String'; ofType: null; }; defaultValue: null }, { name: 'password'; type: { kind: 'SCALAR'; name: 'String'; ofType: null; }; defaultValue: null }, { name: 'email'; type: { kind: 'SCALAR'; name: 'String'; ofType: null; }; defaultValue: null }, { name: 'role'; type: { kind: 'SCALAR'; name: 'String'; ofType: null; }; defaultValue: null }]; };
    'Upload': unknown;
    'UploadResponse': { kind: 'OBJECT'; name: 'UploadResponse'; fields: { 'is_success': { name: 'is_success'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'Boolean'; ofType: null; }; } }; 's3_url': { name: 's3_url'; type: { kind: 'SCALAR'; name: 'String'; ofType: null; } }; }; };
    'User': { kind: 'OBJECT'; name: 'User'; fields: { 'created_at': { name: 'created_at'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'DateTime'; ofType: null; }; } }; 'email': { name: 'email'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'String'; ofType: null; }; } }; 'id': { name: 'id'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'ObjectID'; ofType: null; }; } }; 'is_active': { name: 'is_active'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'Boolean'; ofType: null; }; } }; 'role': { name: 'role'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'String'; ofType: null; }; } }; 'updated_at': { name: 'updated_at'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'DateTime'; ofType: null; }; } }; 'username': { name: 'username'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'String'; ofType: null; }; } }; }; };
    'Workspace': { kind: 'OBJECT'; name: 'Workspace'; fields: { 'collaborators': { name: 'collaborators'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'LIST'; name: never; ofType: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'User'; ofType: null; }; }; }; } }; 'created_at': { name: 'created_at'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'DateTime'; ofType: null; }; } }; 'excalidraw_objects': { name: 'excalidraw_objects'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'LIST'; name: never; ofType: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'Any'; ofType: null; }; }; }; } }; 'id': { name: 'id'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'ObjectID'; ofType: null; }; } }; 'is_active': { name: 'is_active'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'Boolean'; ofType: null; }; } }; 'is_public': { name: 'is_public'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'Boolean'; ofType: null; }; } }; 'name': { name: 'name'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'String'; ofType: null; }; } }; 'owner': { name: 'owner'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'OBJECT'; name: 'User'; ofType: null; }; } }; 'updated_at': { name: 'updated_at'; type: { kind: 'NON_NULL'; name: never; ofType: { kind: 'SCALAR'; name: 'DateTime'; ofType: null; }; } }; }; };
};

/** An IntrospectionQuery representation of your schema.
 *
 * @remarks
 * This is an introspection of your schema saved as a file by GraphQLSP.
 * It will automatically be used by `gql.tada` to infer the types of your GraphQL documents.
 * If you need to reuse this data or update your `scalars`, update `tadaOutputLocation` to
 * instead save to a .ts instead of a .d.ts file.
 */
export type introspection = {
  name: never;
  query: 'Query';
  mutation: 'Mutation';
  subscription: 'Subscription';
  types: introspection_types;
};

import * as gqlTada from 'gql.tada';

declare module 'gql.tada' {
  interface setupSchema {
    introspection: introspection
  }
}