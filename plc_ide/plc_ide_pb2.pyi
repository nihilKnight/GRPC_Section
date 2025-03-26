from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class LogSubscription(_message.Message):
    __slots__ = ("start_time", "end_time", "level_filter", "source_filter")
    class LogLevel(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        DEBUG: _ClassVar[LogSubscription.LogLevel]
        INFO: _ClassVar[LogSubscription.LogLevel]
        WARNING: _ClassVar[LogSubscription.LogLevel]
        ERROR: _ClassVar[LogSubscription.LogLevel]
        CRITICAL: _ClassVar[LogSubscription.LogLevel]
    DEBUG: LogSubscription.LogLevel
    INFO: LogSubscription.LogLevel
    WARNING: LogSubscription.LogLevel
    ERROR: LogSubscription.LogLevel
    CRITICAL: LogSubscription.LogLevel
    START_TIME_FIELD_NUMBER: _ClassVar[int]
    END_TIME_FIELD_NUMBER: _ClassVar[int]
    LEVEL_FILTER_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FILTER_FIELD_NUMBER: _ClassVar[int]
    start_time: _timestamp_pb2.Timestamp
    end_time: _timestamp_pb2.Timestamp
    level_filter: LogSubscription.LogLevel
    source_filter: str
    def __init__(self, start_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., end_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., level_filter: _Optional[_Union[LogSubscription.LogLevel, str]] = ..., source_filter: _Optional[str] = ...) -> None: ...

class LogEntry(_message.Message):
    __slots__ = ("timestamp", "level", "source", "message", "context")
    class ContextEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    LEVEL_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    CONTEXT_FIELD_NUMBER: _ClassVar[int]
    timestamp: _timestamp_pb2.Timestamp
    level: LogSubscription.LogLevel
    source: str
    message: str
    context: _containers.ScalarMap[str, str]
    def __init__(self, timestamp: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., level: _Optional[_Union[LogSubscription.LogLevel, str]] = ..., source: _Optional[str] = ..., message: _Optional[str] = ..., context: _Optional[_Mapping[str, str]] = ...) -> None: ...

class LogExportRequest(_message.Message):
    __slots__ = ("format", "chunk_size")
    FORMAT_FIELD_NUMBER: _ClassVar[int]
    CHUNK_SIZE_FIELD_NUMBER: _ClassVar[int]
    format: str
    chunk_size: int
    def __init__(self, format: _Optional[str] = ..., chunk_size: _Optional[int] = ...) -> None: ...

class LogChunk(_message.Message):
    __slots__ = ("data", "sequence", "checksum")
    DATA_FIELD_NUMBER: _ClassVar[int]
    SEQUENCE_FIELD_NUMBER: _ClassVar[int]
    CHECKSUM_FIELD_NUMBER: _ClassVar[int]
    data: bytes
    sequence: int
    checksum: str
    def __init__(self, data: _Optional[bytes] = ..., sequence: _Optional[int] = ..., checksum: _Optional[str] = ...) -> None: ...

class ConfigChunk(_message.Message):
    __slots__ = ("content", "total_chunks", "current_chunk", "config_hash")
    CONTENT_FIELD_NUMBER: _ClassVar[int]
    TOTAL_CHUNKS_FIELD_NUMBER: _ClassVar[int]
    CURRENT_CHUNK_FIELD_NUMBER: _ClassVar[int]
    CONFIG_HASH_FIELD_NUMBER: _ClassVar[int]
    content: bytes
    total_chunks: int
    current_chunk: int
    config_hash: str
    def __init__(self, content: _Optional[bytes] = ..., total_chunks: _Optional[int] = ..., current_chunk: _Optional[int] = ..., config_hash: _Optional[str] = ...) -> None: ...

class ImportStatus(_message.Message):
    __slots__ = ("state", "progress", "error_detail")
    class State(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        RECEIVING: _ClassVar[ImportStatus.State]
        VALIDATING: _ClassVar[ImportStatus.State]
        APPLIED: _ClassVar[ImportStatus.State]
        FAILED: _ClassVar[ImportStatus.State]
    RECEIVING: ImportStatus.State
    VALIDATING: ImportStatus.State
    APPLIED: ImportStatus.State
    FAILED: ImportStatus.State
    STATE_FIELD_NUMBER: _ClassVar[int]
    PROGRESS_FIELD_NUMBER: _ClassVar[int]
    ERROR_DETAIL_FIELD_NUMBER: _ClassVar[int]
    state: ImportStatus.State
    progress: float
    error_detail: str
    def __init__(self, state: _Optional[_Union[ImportStatus.State, str]] = ..., progress: _Optional[float] = ..., error_detail: _Optional[str] = ...) -> None: ...

class ExportRequest(_message.Message):
    __slots__ = ("version_id", "include_assets")
    VERSION_ID_FIELD_NUMBER: _ClassVar[int]
    INCLUDE_ASSETS_FIELD_NUMBER: _ClassVar[int]
    version_id: str
    include_assets: bool
    def __init__(self, version_id: _Optional[str] = ..., include_assets: bool = ...) -> None: ...

class VersionQuery(_message.Message):
    __slots__ = ("project_id", "page_size", "page_token")
    PROJECT_ID_FIELD_NUMBER: _ClassVar[int]
    PAGE_SIZE_FIELD_NUMBER: _ClassVar[int]
    PAGE_TOKEN_FIELD_NUMBER: _ClassVar[int]
    project_id: str
    page_size: int
    page_token: str
    def __init__(self, project_id: _Optional[str] = ..., page_size: _Optional[int] = ..., page_token: _Optional[str] = ...) -> None: ...

class ProjectVersions(_message.Message):
    __slots__ = ("versions",)
    class VersionInfo(_message.Message):
        __slots__ = ("id", "created_at", "author", "comment")
        ID_FIELD_NUMBER: _ClassVar[int]
        CREATED_AT_FIELD_NUMBER: _ClassVar[int]
        AUTHOR_FIELD_NUMBER: _ClassVar[int]
        COMMENT_FIELD_NUMBER: _ClassVar[int]
        id: str
        created_at: _timestamp_pb2.Timestamp
        author: str
        comment: str
        def __init__(self, id: _Optional[str] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., author: _Optional[str] = ..., comment: _Optional[str] = ...) -> None: ...
    VERSIONS_FIELD_NUMBER: _ClassVar[int]
    versions: _containers.RepeatedCompositeFieldContainer[ProjectVersions.VersionInfo]
    def __init__(self, versions: _Optional[_Iterable[_Union[ProjectVersions.VersionInfo, _Mapping]]] = ...) -> None: ...

class ProjectMetadata(_message.Message):
    __slots__ = ("project_id", "name", "description", "tags", "dependencies")
    class TagsEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    PROJECT_ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    DEPENDENCIES_FIELD_NUMBER: _ClassVar[int]
    project_id: str
    name: str
    description: str
    tags: _containers.ScalarMap[str, str]
    dependencies: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, project_id: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., tags: _Optional[_Mapping[str, str]] = ..., dependencies: _Optional[_Iterable[str]] = ...) -> None: ...

class ProjectCreationResponse(_message.Message):
    __slots__ = ("project_id", "created_at", "initial_version")
    PROJECT_ID_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    INITIAL_VERSION_FIELD_NUMBER: _ClassVar[int]
    project_id: str
    created_at: _timestamp_pb2.Timestamp
    initial_version: str
    def __init__(self, project_id: _Optional[str] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., initial_version: _Optional[str] = ...) -> None: ...

class ProjectIdentifier(_message.Message):
    __slots__ = ("project_id", "force_delete")
    PROJECT_ID_FIELD_NUMBER: _ClassVar[int]
    FORCE_DELETE_FIELD_NUMBER: _ClassVar[int]
    project_id: str
    force_delete: bool
    def __init__(self, project_id: _Optional[str] = ..., force_delete: bool = ...) -> None: ...

class ProjectFilter(_message.Message):
    __slots__ = ("name_pattern", "tags")
    NAME_PATTERN_FIELD_NUMBER: _ClassVar[int]
    TAGS_FIELD_NUMBER: _ClassVar[int]
    name_pattern: str
    tags: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, name_pattern: _Optional[str] = ..., tags: _Optional[_Iterable[str]] = ...) -> None: ...

class ProjectList(_message.Message):
    __slots__ = ("projects",)
    class ProjectSummary(_message.Message):
        __slots__ = ("project_id", "name", "last_modified", "version_count")
        PROJECT_ID_FIELD_NUMBER: _ClassVar[int]
        NAME_FIELD_NUMBER: _ClassVar[int]
        LAST_MODIFIED_FIELD_NUMBER: _ClassVar[int]
        VERSION_COUNT_FIELD_NUMBER: _ClassVar[int]
        project_id: str
        name: str
        last_modified: _timestamp_pb2.Timestamp
        version_count: int
        def __init__(self, project_id: _Optional[str] = ..., name: _Optional[str] = ..., last_modified: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., version_count: _Optional[int] = ...) -> None: ...
    PROJECTS_FIELD_NUMBER: _ClassVar[int]
    projects: _containers.RepeatedCompositeFieldContainer[ProjectList.ProjectSummary]
    def __init__(self, projects: _Optional[_Iterable[_Union[ProjectList.ProjectSummary, _Mapping]]] = ...) -> None: ...

class OperationStatus(_message.Message):
    __slots__ = ("success", "error_code", "operation_time")
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    ERROR_CODE_FIELD_NUMBER: _ClassVar[int]
    OPERATION_TIME_FIELD_NUMBER: _ClassVar[int]
    success: bool
    error_code: str
    operation_time: _timestamp_pb2.Timestamp
    def __init__(self, success: bool = ..., error_code: _Optional[str] = ..., operation_time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
