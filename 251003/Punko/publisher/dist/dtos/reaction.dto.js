var __esDecorate = (this && this.__esDecorate) || function (ctor, descriptorIn, decorators, contextIn, initializers, extraInitializers) {
    function accept(f) { if (f !== void 0 && typeof f !== "function") throw new TypeError("Function expected"); return f; }
    var kind = contextIn.kind, key = kind === "getter" ? "get" : kind === "setter" ? "set" : "value";
    var target = !descriptorIn && ctor ? contextIn["static"] ? ctor : ctor.prototype : null;
    var descriptor = descriptorIn || (target ? Object.getOwnPropertyDescriptor(target, contextIn.name) : {});
    var _, done = false;
    for (var i = decorators.length - 1; i >= 0; i--) {
        var context = {};
        for (var p in contextIn) context[p] = p === "access" ? {} : contextIn[p];
        for (var p in contextIn.access) context.access[p] = contextIn.access[p];
        context.addInitializer = function (f) { if (done) throw new TypeError("Cannot add initializers after decoration has completed"); extraInitializers.push(accept(f || null)); };
        var result = (0, decorators[i])(kind === "accessor" ? { get: descriptor.get, set: descriptor.set } : descriptor[key], context);
        if (kind === "accessor") {
            if (result === void 0) continue;
            if (result === null || typeof result !== "object") throw new TypeError("Object expected");
            if (_ = accept(result.get)) descriptor.get = _;
            if (_ = accept(result.set)) descriptor.set = _;
            if (_ = accept(result.init)) initializers.unshift(_);
        }
        else if (_ = accept(result)) {
            if (kind === "field") initializers.unshift(_);
            else descriptor[key] = _;
        }
    }
    if (target) Object.defineProperty(target, contextIn.name, descriptor);
    done = true;
};
var __runInitializers = (this && this.__runInitializers) || function (thisArg, initializers, value) {
    var useValue = arguments.length > 2;
    for (var i = 0; i < initializers.length; i++) {
        value = useValue ? initializers[i].call(thisArg, value) : initializers[i].call(thisArg);
    }
    return useValue ? value : void 0;
};
import { IsString, IsNotEmpty } from 'class-validator';
let ReactionRequestTo = (() => {
    var _a;
    let _content_decorators;
    let _content_initializers = [];
    let _content_extraInitializers = [];
    let _issueId_decorators;
    let _issueId_initializers = [];
    let _issueId_extraInitializers = [];
    return _a = class ReactionRequestTo {
            constructor() {
                this.content = __runInitializers(this, _content_initializers, void 0);
                this.issueId = (__runInitializers(this, _content_extraInitializers), __runInitializers(this, _issueId_initializers, void 0));
                __runInitializers(this, _issueId_extraInitializers);
            }
        },
        (() => {
            const _metadata = typeof Symbol === "function" && Symbol.metadata ? Object.create(null) : void 0;
            _content_decorators = [IsString(), IsNotEmpty()];
            _issueId_decorators = [IsNotEmpty()];
            __esDecorate(null, null, _content_decorators, { kind: "field", name: "content", static: false, private: false, access: { has: obj => "content" in obj, get: obj => obj.content, set: (obj, value) => { obj.content = value; } }, metadata: _metadata }, _content_initializers, _content_extraInitializers);
            __esDecorate(null, null, _issueId_decorators, { kind: "field", name: "issueId", static: false, private: false, access: { has: obj => "issueId" in obj, get: obj => obj.issueId, set: (obj, value) => { obj.issueId = value; } }, metadata: _metadata }, _issueId_initializers, _issueId_extraInitializers);
            if (_metadata) Object.defineProperty(_a, Symbol.metadata, { enumerable: true, configurable: true, writable: true, value: _metadata });
        })(),
        _a;
})();
export { ReactionRequestTo };
let ReactionResponseTo = (() => {
    var _a;
    let _content_decorators;
    let _content_initializers = [];
    let _content_extraInitializers = [];
    let _issueId_decorators;
    let _issueId_initializers = [];
    let _issueId_extraInitializers = [];
    return _a = class ReactionResponseTo {
            constructor() {
                this.content = __runInitializers(this, _content_initializers, void 0);
                this.issueId = (__runInitializers(this, _content_extraInitializers), __runInitializers(this, _issueId_initializers, void 0));
                __runInitializers(this, _issueId_extraInitializers);
            }
        },
        (() => {
            const _metadata = typeof Symbol === "function" && Symbol.metadata ? Object.create(null) : void 0;
            _content_decorators = [IsString(), IsNotEmpty()];
            _issueId_decorators = [IsNotEmpty()];
            __esDecorate(null, null, _content_decorators, { kind: "field", name: "content", static: false, private: false, access: { has: obj => "content" in obj, get: obj => obj.content, set: (obj, value) => { obj.content = value; } }, metadata: _metadata }, _content_initializers, _content_extraInitializers);
            __esDecorate(null, null, _issueId_decorators, { kind: "field", name: "issueId", static: false, private: false, access: { has: obj => "issueId" in obj, get: obj => obj.issueId, set: (obj, value) => { obj.issueId = value; } }, metadata: _metadata }, _issueId_initializers, _issueId_extraInitializers);
            if (_metadata) Object.defineProperty(_a, Symbol.metadata, { enumerable: true, configurable: true, writable: true, value: _metadata });
        })(),
        _a;
})();
export { ReactionResponseTo };
